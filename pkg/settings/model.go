package settings

import (
	"github.com/MontFerret/besynes/pkg/common/dal"
)

type (
	Settings struct {
		CDPAddress string `json:"cdpAddress"`
	}

	SettingsDetails struct {
		dal.Metadata
		Settings
	}
)

func NewDefault() Settings {
	return Settings{
		CDPAddress: "http://127.0.0.1:9222",
	}
}
