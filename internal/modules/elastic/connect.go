package elastic

import (
	"calend/internal/modules/config"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

func connectElastic(cfg config.Config, logger *zap.Logger) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%s", cfg.ElasticHost, cfg.ElasticPort)),
		elastic.SetSniff(false),
		elastic.SetTraceLog(newLogger(logger)),
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к Elastic: %w", err)
	}
	return client, nil
}
