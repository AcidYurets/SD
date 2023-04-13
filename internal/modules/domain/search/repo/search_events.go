package repo

import (
	"calend/internal/modules/db"
	"calend/internal/modules/db/ent"
	event_ent "calend/internal/modules/db/ent/event"
	"calend/internal/modules/db/ent/predicate"
	ev_dto "calend/internal/modules/domain/event/dto"
	"calend/internal/modules/domain/event/repo"
	"calend/internal/modules/domain/search/dto"
	search_pkg "calend/internal/pkg/search/engine/db"
	"calend/internal/pkg/search/engine/db/types"
	"calend/internal/pkg/search/paginate"
	"context"
	"strings"
)

type SearchRepo struct {
	client *ent.Client
}

func NewSearchRepo(client *ent.Client) *SearchRepo {
	return &SearchRepo{
		client: client,
	}
}

func (r *SearchRepo) SearchEvents(ctx context.Context, searchRequest *dto.EventSearchRequest) (ev_dto.Events, error) {
	filterPredicates := r.buildEventFilters(searchRequest.Filter)
	sortPredicates := r.buildEventSorts(searchRequest.Sort)
	limit, offset := paginate.BuildPaginate(searchRequest.Paginate)

	events, err := r.client.Event.Query().
		Where(filterPredicates...).
		Order(sortPredicates...).
		Limit(limit).
		Offset(offset).
		WithInvitations().
		All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return repo.ToEventDTOs(events), nil
}

func (r *SearchRepo) buildEventFilters(f *dto.EventFilter) []predicate.Event {
	if f == nil {
		return nil
	}

	ftsSearch := []string{
		"name",
		"description",
		"type",
		"users.login",
		"tags.name",
	}

	// Составляем карту для получения полей из связанных сущностей
	wrappersMap := map[string]types.Wrapper{
		"users.login": func(p types.Predicate) types.Predicate {
			return types.Predicate(event_ent.HasCreatorWith(predicate.User(p)))
		},
		"tags.name": func(p types.Predicate) types.Predicate {
			return types.Predicate(event_ent.HasTagsWith(predicate.Tag(p)))
		},
	}

	builder := search_pkg.NewQueryBuilder(wrappersMap)

	builder.AddField("timestamp", f.Timestamp)
	builder.AddField("name", f.Name)
	builder.AddField("description", f.Description)
	builder.AddField("type", f.Type)
	builder.AddField("is_whole_day", f.IsWholeDay)
	builder.AddField("creator_uuid", f.CreatorUuid)
	builder.AddField("users.login", f.CreatorLogin)
	builder.AddField("tags.name", f.TagName)
	builder.AddField(strings.Join(ftsSearch, " "), f.FTSearchStr)

	predicates := builder.Build()

	var eventPreds []predicate.Event
	for _, p := range predicates {
		eventPreds = append(eventPreds, predicate.Event(p))
	}

	return eventPreds
}

func (r *SearchRepo) buildEventSorts(f *dto.EventSort) []event_ent.Order {
	if f == nil {
		return nil
	}

	// Составляем карту для получения полей из связанных сущностей
	wrappersMap := map[string]types.Wrapper{
		"users.login": func(p types.Predicate) types.Predicate {
			return types.Predicate(event_ent.HasCreatorWith(predicate.User(p)))
		},
	}

	builder := search_pkg.NewSortBuilder(wrappersMap)

	builder.AddField("timestamp", f.Timestamp)
	builder.AddField("name", f.Name)
	builder.AddField("description", f.Description)
	builder.AddField("type", f.Type)
	builder.AddField("users.login", f.CreatorLogin)

	predicates := builder.Build()

	var eventSorts []event_ent.Order
	for _, p := range predicates {
		eventSorts = append(eventSorts, event_ent.Order(p))
	}

	return eventSorts
}
