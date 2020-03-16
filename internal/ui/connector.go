package ui

import (
	"github.com/rs/zerolog"
	"runtime"

	"github.com/pkg/errors"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"

	"github.com/MontFerret/besynes/internal/ui/bridges"
	"github.com/MontFerret/besynes/internal/ui/controllers"
)

type (
	Handler func() (interface{}, error)

	BridgeConnector struct {
		logger   zerolog.Logger
		jsEngine *qml.QJSEngine
		async    *bridges.AsyncHelper
	}
)

func NewBridgeConnector(
	logger zerolog.Logger,
	jsEngine *qml.QJSEngine,
	async *bridges.AsyncHelper,
) *BridgeConnector {
	return &BridgeConnector{logger, jsEngine, async}
}

func (bc *BridgeConnector) ConnectSettings(bridge *bridges.Settings, ctl *controllers.Settings) {
	bridge.ConnectGet(func(callback *qml.QJSValue) {
		bc.run(func() (interface{}, error) {
			return ctl.Get()
		}, callback)
	})

	bridge.ConnectSave(func(values *core.QJsonObject, callback *qml.QJSValue) {
		bc.run(func() (interface{}, error) {
			return ctl.Save(values)
		}, callback)
	})
}

func (bc *BridgeConnector) ConnectExecution(bridge *bridges.Execution, ctl *controllers.Execution) {
	bridge.ConnectExecute(func(query *core.QJsonObject, callback *qml.QJSValue) {
		bc.run(func() (interface{}, error) {
			return ctl.Execute(query)
		}, callback)
	})
}

func (bc *BridgeConnector) run(handler Handler, callback *qml.QJSValue) {
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

				bc.logger.Error().
					Timestamp().
					Err(err).
					Str("stack", string(b)).
					Msg("Panic")
			}
		}()

		out, err := handler()

		// Invoke the callback on the main thread
		bc.async.Run(func() {
			if err != nil {
				callback.Call([]*qml.QJSValue{
					qml.NewQJSValue8(err.Error()),
				})

				return
			}

			callback.Call([]*qml.QJSValue{
				qml.NewQJSValue8(""),
				bc.jsEngine.NewGoType(out),
			})
		})
	}()
}
