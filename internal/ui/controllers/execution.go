package controllers

import (
	"context"
	"github.com/pkg/errors"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"runtime"

	"github.com/rs/zerolog"

	"github.com/MontFerret/besynes/internal/ui/bridges"
	"github.com/MontFerret/besynes/pkg/execution"
)

type Execution struct {
	logger  zerolog.Logger
	service *execution.Service
}

func NewExecution(bridge *bridges.Execution, logger zerolog.Logger, service *execution.Service) *Execution {
	ctl := &Execution{
		logger:  logger,
		service: service,
	}

	bridge.ConnectExecute(ctl.execute)

	return ctl
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

		text := query.Value("text")

		data, err := ctl.service.Execute(context.Background(), execution.Query{
			Text:       text.ToString(),
			Params:     nil,
			CDPAddress: "127.0.0.1:9222",
		})

		args := make([]*qml.QJSValue, 0, 2)

		if err != nil {
			args = append(args, qml.NewQJSValue10(err.Error()))
		} else {
			args = append(args, qml.NewQJSValue10(""), qml.NewQJSValue10(string(data)))
		}

		callback.Call(args)
	}()
}
