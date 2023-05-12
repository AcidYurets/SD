package elastic

import (
	"calend/internal/modules/config"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"

	_ "calend/internal/modules/db/ent/runtime"
)

func NewClient(cfg config.Config, logger *zap.Logger) (*elastic.Client, error) {
	client, err := connectElastic(cfg, logger)
	if err != nil {
		return nil, err
	}

	return client, nil
}
