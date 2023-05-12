package reindex

import (
	"calend/internal/modules/elastic/index"
	"calend/internal/modules/logger"
	"calend/internal/utils/my_cast"
	"context"
	"fmt"
	"go.uber.org/zap"
	"sync/atomic"
	"time"
)

func ByIds(ctx context.Context, config *Config) (Stats, error) {
	log := logger.GetFromCtx(ctx).Named(config.Config.Index).With(zap.String("index", config.Config.Index))
	log.Debug("Начало индексации")

	stats, err := byIds(ctx, config)
	if err != nil {
		log.Error("Индексация завершена с ошибкой", zap.Any("stats", stats), zap.Error(err))
		return stats, err
	}
	log.Debug("Индексация завершена", zap.Any("stats", stats))

	return stats, nil
}

func byIds(ctx context.Context, config *Config) (stats Stats, err error) {
	bulkCount := config.BulkCount

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

	idx := index.NewElasticIndex(config.Config)
	offset := 0
	isDone := false

	for !isDone {
		select {
		case <-ctx.Done():
			return getStats(), ctx.Err()
		default:
			// достаём данные из источника
			var fetchData interface{}
			var loopErr error

			if fetchData, loopErr = config.FetchData(ctx, offset, bulkCount); loopErr != nil {
				err = fmt.Errorf("не удалось получить данные: %w", loopErr)
				isDone = true
				break
			}

			data, loopErr := my_cast.ToInterfaceSliceE(fetchData)
			if loopErr != nil {
				err = fmt.Errorf("не удалось преобразовать данные в массив интерфейсов: %w", loopErr)
				isDone = true
				break
			}

			if len(data) == 0 {
				isDone = true
				break
			}

			offset += bulkCount

			if len(data) < bulkCount {
				isDone = true
			}

			saveStat, loopErr := (&indexingData{}).
				Data(fetchData).
				Index(idx).
				Construct(config.BatchConstruct).
				Do(ctx)

			if loopErr != nil {
				err = fmt.Errorf("не удалось сохранить данные: %w", loopErr)
				break
			}

			atomic.AddUint32(&stat.Total, uint32(saveStat.Total))
			atomic.AddUint32(&stat.Created, uint32(saveStat.Created))
			atomic.AddUint32(&stat.Updated, uint32(saveStat.Updated))
			atomic.AddUint32(&stat.Indexed, uint32(saveStat.Indexed))

		}
	}
	if err != nil {
		return getStats(), err
	}

	if config.RemoveData != nil {
		data, err := config.RemoveData(ctx)
		if err != nil {
			return getStats(), fmt.Errorf("не удалось удалить данные: %w", err)
		}

		if data != nil {
			deleted, err := idx.DeleteBulk(ctx, data, true)
			if err != nil {
				return getStats(), fmt.Errorf("не удалось удалить данные: %w", err)
			}
			atomic.AddUint32(&stat.Removed, uint32(deleted))
			atomic.AddUint32(&stat.Total, uint32(deleted))
		}
	}

	stats = getStats()

	return stats, nil
}

type indexingData struct {
	idx       index.IElasticIndex
	construct func(ctx context.Context, data interface{}) error
	data      interface{}
}

func (i *indexingData) Index(idx index.IElasticIndex) *indexingData {

	i.idx = idx

	return i
}

func (i *indexingData) Construct(fn func(ctx context.Context, data interface{}) error) *indexingData {

	i.construct = fn
	return i
}

func (i *indexingData) Data(data interface{}) *indexingData {

	if data == nil {
		return i
	}

	i.data = data

	return i
}

func (i *indexingData) Do(ctx context.Context) (stats index.BulkStats, err error) {

	if i.construct != nil {
		if err := i.construct(ctx, i.data); err != nil {
			return stats, err
		}
	}

	// сохраняем
	saveStat, err := i.idx.SaveBulk(ctx, i.data, true)
	if err != nil {
		return saveStat, err
	}

	return saveStat, nil
}
