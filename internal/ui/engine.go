package ui

import (
	"github.com/MontFerret/besynes/internal/ui/controllers"
	"os"

	"github.com/rs/zerolog"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"

	"github.com/MontFerret/besynes/internal/ui/bridges"
	"github.com/MontFerret/besynes/pkg/execution"
)

type Engine struct {
	logger   zerolog.Logger
	executor *execution.Service
	window   *gui.QGuiApplication
	app      *qml.QQmlApplicationEngine
}

func New(
	logger zerolog.Logger,
	executor *execution.Service,
) (*Engine, error) {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	return &Engine{
		logger:   logger,
		executor: executor,
		window:   gui.NewQGuiApplication(len(os.Args), os.Args),
		app:      qml.NewQQmlApplicationEngine(nil),
	}, nil
}

func (e *Engine) Run() error {
	execBridge := bridges.NewExecution(nil)

	_ = controllers.NewExecution(execBridge, e.logger, e.executor)

	e.app.RootContext().SetContextProperty("queryApi", execBridge)
	e.app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	e.window.Exec()

	return nil
}
