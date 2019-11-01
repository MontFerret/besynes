package dal

import "errors"

var (
	ErrBeginTransaction = errors.New("unable to begin transaction")
)
