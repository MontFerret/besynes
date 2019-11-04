package controllers

import (
	"github.com/rs/zerolog"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"

	"github.com/MontFerret/besynes/internal/ui/bridges"
	"github.com/MontFerret/besynes/pkg/settings"
)

type Settings struct {
	logger   zerolog.Logger
	jsEngine *qml.QJSEngine
	service  *settings.Service
	async    *bridges.AsyncHelper
	bridge   *bridges.Settings
}

func NewSettings(
	logger zerolog.Logger,
	jsEngine *qml.QJSEngine,
	service *settings.Service,
) *Settings {
	return &Settings{
		logger:   logger,
		jsEngine: jsEngine,
		service:  service,
	}
}

func (ctl *Settings) Connect(async *bridges.AsyncHelper, bridge *bridges.Settings) {
	if ctl.bridge != nil {
		ctl.bridge.DisconnectGet()
		ctl.bridge.DisconnectSave()
	}

	ctl.async = async
	ctl.bridge = bridge
	ctl.bridge.ConnectGet(ctl.get)
}

func (ctl *Settings) get(callback *qml.QJSValue) {
	// No need to start a new goroutine, service uses cached value
	values := ctl.service.Get()

	qvar := core.NewQVariant1(values)
	jserr := qml.NewQJSValue8("")
	jsv := ctl.jsEngine.ToScriptValue(qvar)

	callback.Call([]*qml.QJSValue{jserr, jsv})
}
