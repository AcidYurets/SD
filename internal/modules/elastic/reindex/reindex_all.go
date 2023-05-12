package reindex

import (
	"calend/internal/modules/elastic/index"
	"calend/internal/modules/logger"
	"calend/internal/utils/my_cast"
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

func All(ctx context.Context, config *Config) (stats Stats, err error) {
	log := logger.GetFromCtx(ctx).Named(config.Config.Index).With(zap.String("index", config.Config.Index))
	log.Info("Начало индексации")

	bulkCount := config.BulkCount
	offset := 0
	stat := &Stats{
		StartAt: time.Now(),
	}

	getStats := func() Stats {
		return Stats{
			Removed: atomic.LoadUint32(&stat.Removed),
			Total:   atomic.LoadUint32(&stat.Total),
			Created: atomic.LoadUint32(&stat.Created),
			Updated: atomic.LoadUint32(&stat.Updated),
			Indexed: atomic.LoadUint32(&stat.Indexed),
			StartAt: stat.StartAt,
			EndAt:   time.Now(),
		}
	}

	defer func() {
		if r := recover(); r != nil {
			stats = getStats()

			var ok bool
			if err, ok = r.(error); !ok {
				// Set error that will call the global error handler
				err = fmt.Errorf("%v", r)
				log.Error("Произошла ошибка индексации", zap.Error(err), zap.Any("stats", stats))
			}

		}
	}()

	idx := index.NewElasticIndex(config.Config)

	isDone := false
	for !isDone {
		select {
		case <-ctx.Done():
			return getStats(), ctx.Err()
		default:
			// достаём данные из источника

			fetchData, err := config.FetchData(ctx, offset, bulkCount)

			if err != nil {
				return getStats(), fmt.Errorf("не удалось получить данные: %w", err)
			}

			if config.BatchConstruct != nil {
				if err := config.BatchConstruct(ctx, fetchData); err != nil {
					return getStats(), fmt.Errorf("не удалось обработать данные: %w", err)
				}
			}

			data, err := my_cast.ToInterfaceSliceE(fetchData)
			if err != nil {
				return getStats(), fmt.Errorf("не удалось преобразовать данные: %w", err)
			}

			offset += bulkCount

			if len(data) < bulkCount {
				isDone = true
			}

			if len(data) == 0 {
				isDone = true
			} else {
				// сохраняем
				saveStat, err := idx.SaveBulk(ctx, data, true)
				if err != nil {
					return getStats(), err
				}
				atomic.AddUint32(&stat.Total, uint32(saveStat.Total))
				atomic.AddUint32(&stat.Created, uint32(saveStat.Created))
				atomic.AddUint32(&stat.Updated, uint32(saveStat.Updated))
				atomic.AddUint32(&stat.Indexed, uint32(saveStat.Indexed))
				log.Debug(fmt.Sprintf("часть данных индексирована %d/%d", saveStat.Total, offset-bulkCount+len(data)),
					zap.Any("chunk_stats", saveStat),
					zap.Any("total", offset-bulkCount+len(data)),
				)
			}

		}
	}
	deleted, err := deleteByIndexedAt(ctx, idx, stat.StartAt)
	if err != nil {
		return getStats(), fmt.Errorf("не удалось удалить данные: %w", err)
	}
	atomic.AddUint32(&stat.Removed, uint32(deleted))
	atomic.AddUint32(&stat.Total, uint32(deleted))

	stats = getStats()

	log.Info("Индексация завершена", zap.Any("stats", stats))

	return stats, nil
}

func deleteByIndexedAt(ctx context.Context, idx index.IElasticIndex, startAt time.Time) (int64, error) {
	generalQ := elastic.NewBoolQuery().Should(
		elastic.NewRangeQuery(index.IndexedAtField).Lt(startAt),
		elastic.NewBoolQuery().MustNot(elastic.NewExistsQuery(index.IndexedAtField)),
	)

	result, err := idx.GetConfig().Client.
		DeleteByQuery(idx.GetConfig().Index).
		Query(generalQ).WaitForCompletion(true).
		Do(ctx)

	if err != nil {
		return 0, err
	}
	return result.Deleted, nil

}
