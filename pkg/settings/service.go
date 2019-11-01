package settings

import (
	"github.com/pkg/errors"
	"sync/atomic"
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

	current := NewDefault()

	if !exists {
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

func (svc *Service) Get() Settings {
	cached := svc.cache.Load().(Settings)

	return cached
}

func (svc *Service) Save(settings Settings) error {
	if err := svc.db.Update(settings); err != nil {
		return err
	}

	svc.cache.Store(settings)

	return nil
}
