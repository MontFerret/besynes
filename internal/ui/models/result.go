package models

import "github.com/therecipe/qt/core"

type Result struct {
	core.QObject

	_ string `property:"data"`
	_ string `property:"error"`
}

func init() {
	Result_QRegisterMetaType()
	Result_QmlRegisterType2("besynes.models.query", 1, 0, "Result")
}
