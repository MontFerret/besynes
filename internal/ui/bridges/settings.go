package bridges

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type Settings struct {
	core.QObject

	_ func(callback *qml.QJSValue)                           `slot:"get"`
	_ func(values *core.QJsonObject, callback *qml.QJSValue) `slot:"save"`
}
