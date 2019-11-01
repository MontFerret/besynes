package internal

import (
	"github.com/MontFerret/besynes/pkg/common"
	"github.com/MontFerret/besynes/pkg/settings"
	"os"
	"path/filepath"

	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/MontFerret/besynes/internal/db"
	"github.com/MontFerret/besynes/internal/ui"
	"github.com/MontFerret/besynes/pkg/execution"
)

const appDirName = ".besynes"

type Application struct {
	ui *ui.Engine
}

func ensureHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", errors.Wrap(err, "get user home directory")
	}

	appDirPath := filepath.Join(homeDir, appDirName)

	exists := true

	_, err = os.Stat(appDirPath)

	if err != nil {
		if !os.IsNotExist(err) {
			return "", errors.Wrap(err, "check directory existence")
		}

		exists = false
	}

	if exists {
		return appDirPath, nil
	}

	err = os.Mkdir(appDirPath, os.ModePerm)

	if err != nil {
		return "", errors.Wrap(err, "create directory")
	}

	return appDirPath, nil
}

func New() (*Application, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	appDirPath, err := ensureHomeDir()

	if err != nil {
		logger.Err(err).Msg(ErrCreateDirectory.Error())

		return nil, ErrCreateDirectory
	}

	dbManager, err := db.New(db.Settings{
		Dir: appDirPath,
	})

	if err != nil {
		logger.Err(err).Msg(ErrOpenDatabase.Error())

		return nil, ErrOpenDatabase
	}

	settingsSvc, err := settings.New(dbManager.SettingsRepository())

	if err != nil {
		logger.Err(err).Msg(common.ErrUnexpected.Error())

		return nil, err
	}

	fc := compiler.New()

	exec, err := execution.NewExecutor(logger, fc)

	if err != nil {
		return nil, errors.Wrap(err, "execution service")
	}

	uiEngine, err := ui.New(logger, settingsSvc, exec)

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
