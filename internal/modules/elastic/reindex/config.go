package reindex

import (
	"calend/internal/modules/elastic/index"
	"context"
	"time"
)

type Stats struct {
	StartAt time.Time
	EndAt   time.Time
	Total   uint32
	Created uint32
	Updated uint32
	Indexed uint32
	Removed uint32
}

type Config struct {
	Config         index.Config
	BulkCount      int
	FetchData      func(ctx context.Context, offset int, bulkCount int) (interface{}, error)
	RemoveData     func(ctx context.Context) (interface{}, error)
	BatchConstruct func(ctx context.Context, data interface{}) error
}
