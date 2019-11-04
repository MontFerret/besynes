package bridges

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type Execution struct {
	core.QObject

	_ func(query *core.QJsonObject, callback *qml.QJSValue) `slot:"execute"`
}
