package controllers

import (
	"github.com/rs/zerolog"
	"github.com/therecipe/qt/core"

	"github.com/MontFerret/besynes/pkg/common"
	"github.com/MontFerret/besynes/pkg/common/dal"
	"github.com/MontFerret/besynes/pkg/settings"
)

type Settings struct {
	logger  zerolog.Logger
	service *settings.Service
}

func NewSettings(
	logger zerolog.Logger,
	service *settings.Service,
) *Settings {
	return &Settings{
		logger:  logger,
		service: service,
	}
}

func (ctl *Settings) Get() (settings.SettingsDetails, error) {
	return ctl.service.Get(), nil
}

func (ctl *Settings) Save(values *core.QJsonObject) (dal.Metadata, error) {
	if !values.Contains("cdpAddress") {

		return dal.Metadata{}, common.Error(common.ErrMissedArgument, "cdpAddress")
	}

	return ctl.service.Save(settings.Settings{
		CDPAddress: values.Value("cdpAddress").ToString(),
	})
}
