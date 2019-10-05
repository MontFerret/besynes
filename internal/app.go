package internal

import (
	"github.com/MontFerret/besynes/internal/ui"
	"github.com/MontFerret/besynes/pkg/execution"
	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"os"
)

type Application struct {
	ui *ui.Engine
}

func New() (*Application, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	fc := compiler.New()

	exec, err := execution.NewService(execution.NewDefaultSettings(), logger, fc)

	if err != nil {
		return nil, errors.Wrap(err, "execution service")
	}

	uiEngine, err := ui.New(logger, exec)

	if err != nil {
		return nil, errors.Wrap(err, "ui engine")
	}

	return &Application{
		ui: uiEngine,
	}, nil
}

func (app *Application) Run() error {
	return app.ui.Run()
}
