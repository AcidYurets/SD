package repo_elastic

import (
	ev_dto "calend/internal/modules/domain/event/dto"
	event_elastic "calend/internal/modules/domain/event/elastic"
	"calend/internal/modules/domain/search/dto"
	elastic2 "calend/internal/modules/elastic"
	"calend/internal/modules/elastic/index"
	search_pkg "calend/internal/pkg/search/engine/elastic"
	"context"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strings"
)

type SearchRepo struct {
	index index.IElasticIndex
}

func NewSearchRepo(client *elastic.Client) *SearchRepo {
	return &SearchRepo{
		index: index.NewElasticIndex(index.Config{
			Client:  client,
			Index:   event_elastic.EventIndexName,
			IdField: "Uuid",
		}),
	}
}

func (r *SearchRepo) SearchEvents(ctx context.Context, searchRequest *dto.EventSearchRequest) (ev_dto.Events, error) {
	var err error
	generalQ := elastic.NewBoolQuery()

	sorter, paginate := (&search_pkg.SearchRequestBuilder{
		Query: generalQ,
		Filter: func() *elastic.BoolQuery {
			return r.buildEventFilters(searchRequest.Filter)
		},
		Sort: func() func(search *elastic.SearchService) *elastic.SearchService {
			return r.buildEventSorts(searchRequest.Sort)
		},
		Paginate: searchRequest.Paginate,
	}).Build()

	searchResult, err := r.index.SearchQuery(generalQ, sorter, paginate).Do(ctx)
	if err != nil {
		return nil, elastic2.WrapElasticError(err)
	}

	var result ev_dto.Events
	for _, item := range searchResult.Each(reflect.TypeOf(&ev_dto.Event{})) {
		result = append(result, item.(*ev_dto.Event))
	}

	return result, nil
}

func (r *SearchRepo) buildEventFilters(f *dto.EventFilter) *elastic.BoolQuery {
	if f == nil {
		return nil
	}

	ftsSearch := []string{
		"Name",
		"Description",
		"Type",
		"Creator.Login",
		"Tags.Name",
	}

	builder := search_pkg.QueryBuilder{}

	builder.AddField("Timestamp", f.Timestamp)
	builder.AddField("Name", f.Name)
	builder.AddField("Description", f.Description)
	builder.AddField("Type", f.Type)
	builder.AddField("IsWholeDay", f.IsWholeDay)
	builder.AddField("CreatorUuid", f.CreatorUuid)
	builder.AddField("Creator.Login", f.CreatorLogin)
	builder.AddField("Tags.Name", f.TagName)
	builder.AddField("Invitations.UserUuid", f.InvitedUserUuid)
	builder.AddField(strings.Join(ftsSearch, " "), f.FTSearchStr)

	return builder.Build()
}

func (r *SearchRepo) buildEventSorts(f *dto.EventSort) func(search *elastic.SearchService) *elastic.SearchService {
	if f == nil {
		return nil
	}

	builder := search_pkg.SortBuilder{}

	builder.AddField("Timestamp", f.Timestamp)
	builder.AddField("Name", f.Name)
	builder.AddField("Description", f.Description)
	builder.AddField("Type", f.Type)
	builder.AddField("Creator.Login", f.CreatorLogin)

	if val := builder.Build(); len(val) > 0 {
		return func(search *elastic.SearchService) *elastic.SearchService {
			return search.SortBy(val...)
		}
	}

	return nil
}
