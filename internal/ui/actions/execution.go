package actions

import (
	"github.com/MontFerret/besynes/internal/ui/models"
	"github.com/therecipe/qt/core"

	"github.com/MontFerret/besynes/pkg/execution"
)

type Executions struct {
	core.QObject

	executor *execution.Service

	_ func(query, addr string) *models.Result `slot:"execute"`
}
