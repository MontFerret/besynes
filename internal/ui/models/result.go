package models

import "github.com/therecipe/qt/core"

type Result struct {
	core.QObject

	_ string `property:"data"`
	_ string `property:"error"`
}

func init() {
	Result_QRegisterMetaType()
}
