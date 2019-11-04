package bridges

import (
	"github.com/therecipe/qt/core"
)

type AsyncHelper struct {
	core.QObject

	_ func(f func()) `slot:"run,auto"`
}

func (*AsyncHelper) run(f func()) {
	f()
}
