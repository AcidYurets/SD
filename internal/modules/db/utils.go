package db

import (
	"calend/internal/models/err_const"
	"calend/internal/modules/db/ent"
	"calend/internal/modules/db/schema"
	"context"
	"errors"
	"fmt"
	"github.com/lib/pq"
)

func WrapError(err error) error {
	defaultError := err

	// Если запись не найдена
	if ent.IsNotFound(err) {
		return err_const.ErrDatabaseRecordNotFound
	}

	// Если нарушено ограничение целостности
	if ent.IsConstraintError(err) {
		// Получаем обернутую ошибку драйвера
		err = errors.Unwrap(err)
		var pqErr *pq.Error
		switch err.(type) {
		case *pq.Error:
			pqErr = err.(*pq.Error)

		default:
			// Пробуем развернуть еще раз
			err = errors.Unwrap(err)
			var ok bool
			pqErr, ok = err.(*pq.Error)
			if !ok {
				return defaultError
			}
		}

		// Обрабатываем ошибку в зависимости от кода
		switch pqErr.Code {
		case "23505": // Нарушение ограничения уникальности
			err = fmt.Errorf("нарушение уникальности: %s", pqErr.Detail)
		case "23503": // Нарушение ограничения внешнего ключа
			err = fmt.Errorf("нарушение внешнего ключа: %s", pqErr.Detail)
		}
		return err
	}

	return defaultError
}

func TruncateAll(client *ent.Client) error {
	_, err := client.Invitation.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	if err != nil {
		return err
	}

	_, err = client.Event.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	if err != nil {
		return err
	}

	_, err = client.Tag.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	if err != nil {
		return err
	}

	_, err = client.AccessRight.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	if err != nil {
		return err
	}

	_, err = client.User.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	if err != nil {
		return err
	}

	return nil
}
