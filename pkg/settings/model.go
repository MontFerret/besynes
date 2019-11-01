package settings

import (
	"github.com/MontFerret/besynes/pkg/common/dal"
	"time"
)

type Settings struct {
	dal.Metadata
	CDPAddress string `json:"cdp_address"`
}

func NewDefault() Settings {
	return Settings{
		Metadata: dal.Metadata{
			CreatedAt: time.Now(),
			UpdateAt:  nil,
		},
		CDPAddress: "http://127.0.0.1:9222",
	}
}
