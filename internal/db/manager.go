package db

import (
	bolt "go.etcd.io/bbolt"
	"path/filepath"
	"time"

	"github.com/MontFerret/besynes/internal/db/repositories"
	"github.com/MontFerret/besynes/pkg/settings"
)

const dbName = "data.db"

type Manager struct {
	db *bolt.DB

	settings *repositories.Settings
}

func New(settings Settings) (*Manager, error) {
	db, err := bolt.Open(filepath.Join(settings.Dir, dbName), 0666, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return nil, err
	}

	settingsRepo, err := repositories.NewSettings(db)

	if err != nil {
		return nil, err
	}

	return &Manager{db, settingsRepo}, nil
}

func (m *Manager) Close() error {
	return m.db.Close()
}

func (m *Manager) SettingsRepository() settings.Repository {
	return m.settings
}
