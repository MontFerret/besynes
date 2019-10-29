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
)

type Execution struct {
	logger   zerolog.Logger
	jsEngine *qml.QJSEngine
	bridge   *bridges.Execution
	service  *execution.Service
}

func NewExecution(logger zerolog.Logger, jsEngine *qml.QJSEngine, service *execution.Service) *Execution {
	return &Execution{
		logger:   logger,
		jsEngine: jsEngine,
		service:  service,
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

		var text string

		if query.Contains("text") {
			text = query.Value("text").ToString()
		}

		out := ctl.service.Execute(context.Background(), execution.Query{
			Text:       text,
			Params:     nil,
			CDPAddress: "http://127.0.0.1:9222",
		})

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

func (ctl *Execution) formatDuration(d time.Duration) string {
	return d.String()
}

func (ctl *Execution) formatSize(sizeInBytes int) string {
	return bytefmt.ByteSize(uint64(sizeInBytes))
}
