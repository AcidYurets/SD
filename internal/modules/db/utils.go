package db

import (
	"calend/internal/models/err_const"
	"calend/internal/modules/db/ent"
	"errors"
	"fmt"
	"github.com/lib/pq"
)

func WrapError(err error) error {
	// Если запись не найдена
	if ent.IsNotFound(err) {
		return err_const.ErrDatabaseRecordNotFound
	}

	// Если нарушено ограничение целостности
	if ent.IsConstraintError(err) {
		// Получаем обернутую ошибку драйвера
		err = errors.Unwrap(err)
		e, ok := err.(*pq.Error)
		if !ok {
			return err
		}

		// Обрабатываем ошибку в зависимости от кода
		switch e.Code {
		case "23505": // Нарушение ограничения уникальности
			err = fmt.Errorf("нарушение уникальности: %s", e.Detail)
		}
		return err
	}

	return err
}
