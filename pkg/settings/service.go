package settings

import (
	"github.com/MontFerret/besynes/pkg/common/dal"
	"github.com/pkg/errors"
	"sync/atomic"
	"time"
)

type Service struct {
	db    Repository
	cache atomic.Value
}

func New(db Repository) (*Service, error) {
	exists, err := db.Exists()

	if err != nil {
		return nil, errors.Wrap(err, "check settings existence")
	}

	var current SettingsDetails

	if !exists {
		current = toDetails(NewDefault())
		err := db.Create(current)

		if err != nil {
			return nil, errors.Wrap(err, "create initial settings")
		}
	} else {
		found, err := db.Get()

		if err != nil {
			return nil, errors.Wrap(err, "restore current settings")
		}

		current = found
	}

	av := atomic.Value{}
	av.Store(current)

	return &Service{db, av}, nil
}

func (svc *Service) Get() (SettingsDetails, error) {
	cached := svc.cache.Load().(SettingsDetails)

	return cached, nil
}

func (svc *Service) Save(settings Settings) (dal.Metadata, error) {
	details := toDetails(settings)

	err := svc.db.Update(details)

	if err != nil {
		return details.Metadata, err
	}

	svc.cache.Store(details)

	return details.Metadata, nil
}

func toDetails(settings Settings) SettingsDetails {
	meta := dal.Metadata{
		UpdateAt: time.Now(),
	}

	return SettingsDetails{
		Metadata: meta,
		Settings: settings,
	}
}
