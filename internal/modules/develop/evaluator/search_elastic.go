package evaluator

import (
	ev_dto "calend/internal/modules/domain/event/dto"
	"calend/internal/modules/domain/search/dto"
	elastic2 "calend/internal/modules/elastic"
	search_pkg "calend/internal/pkg/search/engine/elastic"
	"context"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strings"
)

func (r *Evaluator) searchEventsElastic(ctx context.Context, searchRequest *dto.EventSearchRequest) (ev_dto.Events, error) {
	var err error
	generalQ := elastic.NewBoolQuery()

	sorter, paginate := (&search_pkg.SearchRequestBuilder{
		Query: generalQ,
		Filter: func() *elastic.BoolQuery {
			return r.buildEventFiltersElastic(searchRequest.Filter)
		},
		Sort: func() func(search *elastic.SearchService) *elastic.SearchService {
			return r.buildEventSortsElastic(searchRequest.Sort)
		},
		Paginate: searchRequest.Paginate,
	}).Build()

	searchResult, err := r.elasticIndex.SearchQuery(generalQ, sorter, paginate).Do(ctx)
	if err != nil {
		return nil, elastic2.WrapElasticError(err)
	}

	var result ev_dto.Events
	for _, item := range searchResult.Each(reflect.TypeOf(&ev_dto.Event{})) {
		result = append(result, item.(*ev_dto.Event))
	}

	return result, nil
}

func (r *Evaluator) buildEventFiltersElastic(f *dto.EventFilter) *elastic.BoolQuery {
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

func (r *Evaluator) buildEventSortsElastic(f *dto.EventSort) func(search *elastic.SearchService) *elastic.SearchService {
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
