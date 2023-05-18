package elastic

import (
	"calend/internal/modules/config"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

func connectElastic(cfg config.Config, logger *zap.Logger) (*elastic.Client, error) {
	opts := make([]elastic.ClientOptionFunc, 0)

	opts = append(opts,
		elastic.SetURL(fmt.Sprintf("http://%s:%s", cfg.ElasticHost, cfg.ElasticPort)),
		elastic.SetSniff(false),
	)

	if cfg.TraceElasticQueries {
		opts = append(opts, elastic.SetTraceLog(newLogger(logger)))
	}

	client, err := elastic.NewClient(opts...)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к Elastic: %w", err)
	}
	return client, nil
}
