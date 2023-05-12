package elastic

import (
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func WrapElasticError(err error) error {
	var esErr *elastic.Error

	if errors.As(err, &esErr) {
		wrappedError := err

		for _, c := range esErr.Details.RootCause {
			wrappedError = fmt.Errorf("%w: %s", wrappedError, c.Reason)
		}
		return fmt.Errorf("ошибка поиска в Elastic: %w", wrappedError)
	}

	return err

}
