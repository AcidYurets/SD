package db

import (
	"calend/internal/models/err_const"
	"calend/internal/modules/db/ent"
	"errors"
)

func WrapError(err error) error {
	if errors.Is(err, &ent.NotFoundError{}) {
		return err_const.ErrDatabaseRecordNotFound
	}

	return err
}
