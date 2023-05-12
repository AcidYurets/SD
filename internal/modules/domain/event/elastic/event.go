package elastic

import (
	"calend/internal/modules/domain/event/dto"
	"calend/internal/modules/elastic/index"
	"calend/internal/modules/elastic/reindex"
	"calend/internal/utils/slice"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

const EventIndexName = "events"
const EventMaxBulkSize = 5000

type IEventRepo interface {
	// SubSet выборка с сортировкой, отступом и лимитом
	SubSet(ctx context.Context, offset, limit int, uuids ...string) (dto.Events, error)
}

type EventElasticService struct {
	repo  IEventRepo
	index index.IElasticIndex
}

func NewEventElasticService(client *elastic.Client, repo IEventRepo) *EventElasticService {
	return &EventElasticService{
		repo: repo,
		index: index.NewElasticIndex(index.Config{
			Client:  client,
			Index:   EventIndexName,
			IdField: "Uuid",
		}),
	}
}

// ReindexByUuids переиндексирует по нескольким Uuids
func (s *EventElasticService) ReindexByUuids(ctx context.Context, uuids ...string) (reindex.Stats, error) {
	var indexedUuids []string

	stats, err := reindex.ByIds(ctx, &reindex.Config{
		Config:    s.index.GetConfig(),
		BulkCount: EventMaxBulkSize,
		FetchData: func(ctx context.Context, offset int, bulkCount int) (interface{}, error) {
			employees, err := s.repo.SubSet(ctx, offset, bulkCount, uuids...)
			if err != nil {
				return nil, fmt.Errorf("ошибка выборки по Uuids: %w", err)
			}

			for _, employee := range employees {
				indexedUuids = append(indexedUuids, employee.Uuid)
			}

			return employees, nil
		},
		RemoveData: func(ctx context.Context) (interface{}, error) {
			if len(indexedUuids) == len(uuids) {
				return nil, nil
			}
			var removeUuids []string

			for _, uuid := range uuids {
				if !slice.Contains(indexedUuids, uuid) {
					removeUuids = append(removeUuids, uuid)
				}
			}

			return removeUuids, nil
		},
	})
	if err != nil {
		return stats, fmt.Errorf("ошибка индексации по Uuids: %w", err)
	}

	return stats, nil
}

// ReindexAll переиндексирует все записи
func (s *EventElasticService) ReindexAll(ctx context.Context) (reindex.Stats, error) {
	stats, err := reindex.All(ctx, &reindex.Config{
		Config:    s.index.GetConfig(),
		BulkCount: EventMaxBulkSize,
		FetchData: func(ctx context.Context, offset int, bulkCount int) (interface{}, error) {
			return s.repo.SubSet(ctx, offset, bulkCount)
		},
	})
	if err != nil {
		return stats, fmt.Errorf("ошибка индексации всех записей: %w", err)
	}

	return stats, nil
}
