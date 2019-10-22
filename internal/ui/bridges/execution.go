package bridges

import (
	"github.com/therecipe/qt/core"

	"github.com/MontFerret/besynes/internal/ui/models"
)

type Execution struct {
	core.QObject

	_ func(query, addr string) *models.Result `slot:"execute"`
}
