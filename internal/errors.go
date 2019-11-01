package internal

import "github.com/pkg/errors"

var (
	ErrCreateDirectory = errors.New("unable to create app directory")
	ErrOpenDatabase    = errors.New("unable to open app database")
)
