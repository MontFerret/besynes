package controllers

import (
	"github.com/rs/zerolog"
	"github.com/therecipe/qt/core"

	"github.com/MontFerret/besynes/pkg/common"
	"github.com/MontFerret/besynes/pkg/common/dal"
	"github.com/MontFerret/besynes/pkg/settings"
)

// TODO: Delete me when data serialization gets fixed
type SettingsViewModel struct {
	CDPAddress string `json:"cdpAddress"`
}

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

func (ctl *Settings) Get() (SettingsViewModel, error) {
	out, err := ctl.service.Get()

	if err != nil {
		return SettingsViewModel{}, err
	}

	return SettingsViewModel{
		CDPAddress: out.Settings.CDPAddress,
	}, nil
}

func (ctl *Settings) Save(values *core.QJsonObject) (dal.Metadata, error) {
	if !values.Contains("cdpAddress") {
		return dal.Metadata{}, common.Error(common.ErrMissedArgument, "cdpAddress")
	}

	return ctl.service.Save(settings.Settings{
		CDPAddress: values.Value("cdpAddress").ToString(),
	})
}
