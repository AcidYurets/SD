package index

import (
	"calend/internal/utils/my_cast"
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/cast"
)

type IElasticIndex interface {
	SearchQuery(query elastic.Query, scopes ...func(search *elastic.SearchService) *elastic.SearchService) *elastic.SearchService

	SaveBulk(ctx context.Context, data interface{}, refresh ...bool) (BulkStats, error)
	DeleteBulk(ctx context.Context, ids interface{}, refresh ...bool) (int, error)

	GetConfig() Config
}

type BulkStats struct {
	Total   int
	Created int
	Updated int
	Indexed int
	Deleted int
	Errored int
}

func NewElasticIndex(config Config) IElasticIndex {

	return &elasticIndex{
		config,
	}
}

type Config struct {
	Client  *elastic.Client
	Index   string
	IdField string
}

type elasticIndex struct {
	Config
}

func (e *elasticIndex) GetConfig() Config {
	return e.Config
}

func (e *elasticIndex) SearchQuery(query elastic.Query, scopes ...func(search *elastic.SearchService) *elastic.SearchService) *elastic.SearchService {
	search := e.Client.Search().Index(e.Index).Query(query)
	for _, scope := range scopes {
		search = scope(search)
	}
	return search

}

func (e *elasticIndex) DeleteBulk(ctx context.Context, ids interface{}, refresh ...bool) (int, error) {
	eids, err := cast.ToSliceE(ids)

	if err != nil {
		return 0, err
	}

	if len(eids) == 0 {
		return 0, nil
	}

	field := e.IdField

	if strings.ToLower(field) == "uuid" {
		field += ".keyword"
	}

	generalQ := elastic.NewBoolQuery().Must(elastic.NewTermsQuery(field, eids...))

	bulkDeleteRequest := e.Client.DeleteByQuery(e.Index).
		Query(generalQ)
	if len(refresh) > 0 && refresh[0] {
		bulkDeleteRequest = bulkDeleteRequest.Refresh("true")
	}

	result, err := bulkDeleteRequest.Do(ctx)
	if err != nil {
		return 0, err
	}

	return int(result.Deleted), err
}

func (e *elasticIndex) SaveBulk(ctx context.Context, data interface{}, refresh ...bool) (BulkStats, error) {
	var (
		parsedSlice map[string]interface{}
		err         error
		isMap       bool
	)

	if parsedSlice, isMap = data.(map[string]interface{}); !isMap {
		parsedSlice, err = my_cast.ToMapByFieldE(e.IdField, data)
		if err != nil {
			return BulkStats{}, err
		}
	}

	bulkSaveRequest := e.Client.Bulk()
	for key, value := range parsedSlice {
		onRequest := elastic.NewBulkIndexRequest().
			Index(e.Index).
			Id(key).
			Type("_doc").
			Doc(IndexedAtData{Data: value})

		bulkSaveRequest = bulkSaveRequest.Add(onRequest)
	}

	if len(refresh) > 0 && refresh[0] {
		bulkSaveRequest = bulkSaveRequest.Refresh("true")
	}

	bulkSaveResponse, err := bulkSaveRequest.Do(ctx)
	if err != nil {
		return BulkStats{}, err
	}

	stat := calculateBulkStats(bulkSaveResponse)

	if stat.Total != len(parsedSlice) {
		return stat, fmt.Errorf("partially saved, expected %d, actual %d", len(parsedSlice), stat.Total)
	}
	return stat, nil
}

func calculateBulkStats(result *elastic.BulkResponse) BulkStats {
	filter := func(str string) int {
		var count int
		for _, value := range result.Succeeded() {
			if strings.EqualFold(value.Result, str) {
				count++
			}
		}
		return count
	}
	stat := BulkStats{
		Created: filter("Created"),
		Updated: filter("Updated"),
		Indexed: filter("Indexed"),
		Deleted: filter("Deleted"),
		Errored: len(result.Failed()),
	}

	stat.Total = stat.Created + stat.Updated + stat.Indexed + stat.Deleted + stat.Errored

	return stat
}
