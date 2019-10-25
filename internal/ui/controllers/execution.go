package controllers

import (
	"context"
	"fmt"
	"github.com/MontFerret/besynes/internal/ui/bridges"
	"runtime"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"

	"github.com/MontFerret/besynes/pkg/execution"
)

type Execution struct {
	logger   zerolog.Logger
	jsEngine *qml.QJSEngine
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
	bridge.ConnectExecute(ctl.execute)
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
			CDPAddress: "127.0.0.1:9222",
		})

		jsv := ctl.jsEngine.NewObject()

		if len(out.Data) > 0 {
			jsv.SetProperty("data", qml.NewQJSValue8(string(out.Data)))
		}

		if out.Error != nil {
			jsv.SetProperty("error", qml.NewQJSValue8(out.Error.Error()))
		}

		jsvStats := ctl.jsEngine.NewObject()
		jsvStats.SetProperty("compilation", qml.NewQJSValue8(fmt.Sprintf("%d ms", out.Stats.Compilation.Milliseconds())))
		jsvStats.SetProperty("runtime", qml.NewQJSValue8(fmt.Sprintf("%d ms", out.Stats.Runtime.Milliseconds())))
		jsvStats.SetProperty("size", qml.NewQJSValue8(fmt.Sprintf("%d kb", len(out.Data)/1000)))

		jsv.SetProperty("stats", jsvStats)

		callback.Call([]*qml.QJSValue{jsv})
	}()
}
