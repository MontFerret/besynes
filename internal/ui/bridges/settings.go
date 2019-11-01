package bridges

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type Settings struct {
	core.QObject

	_ func(f func())                                         `slot:"runOnMainHelper,auto"`
	_ func(callback *qml.QJSValue)                           `slot:"get"`
	_ func(values *core.QJsonObject, callback *qml.QJSValue) `slot:"save"`
}

func (*Settings) runOnMainHelper(f func()) {
	f()
}
