package ui

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"

	"github.com/MontFerret/besynes/internal/ui/bridges"
	"github.com/MontFerret/besynes/internal/ui/controllers"
	"github.com/MontFerret/besynes/pkg/execution"
	"github.com/MontFerret/besynes/pkg/settings"
)

type Engine struct {
	logger   zerolog.Logger
	settings *settings.Service
	executor *execution.Executor
	window   *gui.QGuiApplication
	app      *qml.QQmlApplicationEngine
}

func New(
	logger zerolog.Logger,
	settingsSvc *settings.Service,
	executor *execution.Executor,
) (*Engine, error) {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	return &Engine{
		logger:   logger,
		settings: settingsSvc,
		executor: executor,
		window:   gui.NewQGuiApplication(len(os.Args), os.Args),
		app:      qml.NewQQmlApplicationEngine(nil),
	}, nil
}

func (e *Engine) Run() error {
	execBridge := bridges.NewExecution(nil)
	settingsBridge := bridges.NewSettings(nil)

	connector := NewBridgeConnector(e.logger, e.app.QJSEngine_PTR(), bridges.NewAsyncHelper(nil))
	connector.ConnectExecution(
		execBridge,
		controllers.NewExecution(e.logger, e.settings, e.executor),
	)
	connector.ConnectSettings(
		settingsBridge,
		controllers.NewSettings(e.logger, e.settings),
	)

	e.app.RootContext().SetContextProperty("settingsApi", settingsBridge)
	e.app.RootContext().SetContextProperty("queryApi", execBridge)
	e.app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	e.window.SetAttribute(core.Qt__AA_UseHighDpiPixmaps, true)
	e.window.Exec()

	return nil
}
