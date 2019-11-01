package controllers

import (
	"github.com/rs/zerolog"
	"github.com/therecipe/qt/qml"

	"github.com/MontFerret/besynes/internal/ui/bridges"
	"github.com/MontFerret/besynes/pkg/settings"
)

type Settings struct {
	logger   zerolog.Logger
	jsEngine *qml.QJSEngine
	service  *settings.Service
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

func (ctl *Settings) Connect(bridge *bridges.Settings) {
	if ctl.bridge != nil {
		ctl.bridge.DisconnectGet()
		ctl.bridge.DisconnectSave()
	}

	ctl.bridge = bridge
	ctl.bridge.ConnectGet(ctl.get)
}

func (ctl *Settings) get(callback *qml.QJSValue) {
	// No need to start a new goroutine, service uses cached value
	values := ctl.service.Get()

	jsErr := qml.NewQJSValue8("")

	jsValues := ctl.jsEngine.NewObject()
	jsValues.SetProperty("cdpAddress", qml.NewQJSValue8(values.CDPAddress))

	callback.Call([]*qml.QJSValue{jsErr, jsValues})
}
