package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"runtime"
	"time"

	"github.com/cloudfoundry/bytefmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"

	"github.com/MontFerret/besynes/internal/ui/bridges"
	"github.com/MontFerret/besynes/pkg/execution"
	"github.com/MontFerret/besynes/pkg/settings"
)

var ErrInvalidParams = errors.New("Invalid parameter values. Use valid JSON object.")

type Execution struct {
	logger   zerolog.Logger
	jsEngine *qml.QJSEngine
	bridge   *bridges.Execution
	settings *settings.Service
	executor *execution.Executor
}

func NewExecution(
	logger zerolog.Logger,
	jsEngine *qml.QJSEngine,
	settingsSvc *settings.Service,
	executor *execution.Executor,
) *Execution {
	return &Execution{
		logger:   logger,
		jsEngine: jsEngine,
		settings: settingsSvc,
		executor: executor,
	}
}

func (ctl *Execution) Connect(bridge *bridges.Execution) {
	if ctl.bridge != nil {
		ctl.bridge.DisconnectExecute()
	}

	ctl.bridge = bridge
	ctl.bridge.ConnectExecute(ctl.execute)
}

func (ctl *Execution) execute(query *core.QJsonObject, callback *qml.QJSValue) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				var err error

				// find out exactly what the error was and set err
				switch x := r.(type) {
				case string:
					err = errors.New(x)
				case error:
					err = x
				default:
					err = errors.New("unknown panic")
				}

				b := make([]byte, 0, 20)
				runtime.Stack(b, true)

				ctl.logger.Error().
					Timestamp().
					Err(err).
					Str("stack", string(b)).
					Msg("Panic")
			}
		}()

		q, err := ctl.parseQuery(query)

		if err != nil {
			jsv := ctl.jsEngine.NewObject()
			jsv.SetProperty("error", qml.NewQJSValue8(err.Error()))

			// https://github.com/therecipe/qt/issues/994
			ctl.bridge.RunOnMainHelper(func() {
				callback.Call([]*qml.QJSValue{jsv})
			})

			return
		}

		out := ctl.executor.Execute(context.Background(), q)

		jsv := ctl.jsEngine.NewObject()

		jsv.SetProperty("data", qml.NewQJSValue8(ctl.formatJSON(out.Data)))

		if out.Error != nil {
			jsv.SetProperty("error", qml.NewQJSValue8(out.Error.Error()))
		}

		jsvStats := ctl.jsEngine.NewObject()
		jsvStats.SetProperty("compilation", qml.NewQJSValue8(ctl.formatDuration(out.Stats.Compilation)))
		jsvStats.SetProperty("runtime", qml.NewQJSValue8(ctl.formatDuration(out.Stats.Runtime)))
		jsvStats.SetProperty("size", qml.NewQJSValue8(ctl.formatSize(len(out.Data))))

		jsv.SetProperty("stats", jsvStats)

		// https://github.com/therecipe/qt/issues/994
		ctl.bridge.RunOnMainHelper(func() {
			callback.Call([]*qml.QJSValue{jsv})
		})
	}()
}

func (ctl *Execution) parseQuery(query *core.QJsonObject) (execution.Query, error) {
	var text string
	var params map[string]interface{}

	if query.Contains("text") {
		text = query.Value("text").ToString()
	}

	if query.Contains("params") {
		paramsBytes := []byte(query.Value("params").ToString())

		// Make sure that params string is not empty
		if !ctl.isParamsEmpty(paramsBytes) {
			// Check if it's an object
			if !ctl.isParamsObjectValid(paramsBytes) {
				return execution.Query{}, ErrInvalidParams
			}

			params = make(map[string]interface{})

			err := json.Unmarshal(paramsBytes, &params)

			if err != nil {
				return execution.Query{}, ErrInvalidParams
			}
		}
	}

	opts := ctl.settings.Get()

	return execution.Query{
		Text:       text,
		Params:     params,
		CDPAddress: opts.CDPAddress,
	}, nil
}

func (ctl *Execution) formatJSON(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	var b bytes.Buffer

	err := json.Indent(&b, data, "", "  ")

	if err != nil {
		ctl.logger.Error().Err(err).Msg("failed to format output")

		return string(data)
	}

	return b.String()
}

func (ctl *Execution) isParamsEmpty(text []byte) bool {
	return len(text) <= 2
}

func (ctl *Execution) isParamsObjectValid(text []byte) bool {
	if len(text) < 2 {
		return false
	}

	return string(text[0:1]) == "{" && string(text[len(text)-1:]) == "}"
}

func (ctl *Execution) formatDuration(d time.Duration) string {
	return d.String()
}

func (ctl *Execution) formatSize(sizeInBytes int) string {
	return bytefmt.ByteSize(uint64(sizeInBytes))
}
