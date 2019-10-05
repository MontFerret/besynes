package stores

import (
	"github.com/therecipe/qt/core"
)

type (
	ResultViewModel struct {
		core.QObject
	}


	ExecutionStore struct {
		core.QObject

		_ func(addr string) *ResultViewModel `slot:"execute"`
	}
)
