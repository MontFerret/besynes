package repositories

import (
	"encoding/json"
	"github.com/MontFerret/besynes/pkg/common"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"

	"github.com/MontFerret/besynes/pkg/settings"
)

var (
	settingsBucket = []byte("global_settings")

	settingsBucketKey = []byte("settings")
)

type Settings struct {
	db *bolt.DB
}

func NewSettings(db *bolt.DB) (*Settings, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists(settingsBucket)

		return e
	})

	if err != nil {
		return nil, errors.Wrap(err, "create or open bucket")
	}

	return &Settings{db: db}, nil
}

func (r *Settings) Get() (result settings.Settings, err error) {
	err = r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(settingsBucket)
		b := bucket.Get(settingsBucketKey)

		if b == nil {
			err = common.ErrNotFound

			return nil
		}

		result = settings.Settings{}

		return json.Unmarshal(b, &result)
	})

	return
}

func (r *Settings) Exists() (exists bool, err error) {
	err = r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(settingsBucket)
		b := bucket.Get(settingsBucketKey)

		exists = b != nil

		return nil
	})

	return
}

func (r *Settings) Create(model settings.Settings) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(settingsBucket)
		b := bucket.Get(settingsBucketKey)

		if b != nil {
			return common.ErrNotUnique
		}

		b, err := json.Marshal(model)

		if err != nil {
			return errors.Wrap(err, "model serialization")
		}

		return bucket.Put(settingsBucketKey, b)
	})
}

func (r *Settings) Update(model settings.Settings) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(settingsBucket)
		b := bucket.Get(settingsBucketKey)

		if b == nil {
			return common.ErrNotFound
		}

		b, err := json.Marshal(model)

		if err != nil {
			return errors.Wrap(err, "model serialization")
		}

		return bucket.Put(settingsBucketKey, b)
	})
}
