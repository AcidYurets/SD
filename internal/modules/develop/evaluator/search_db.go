package evaluator

import (
	"calend/internal/modules/db"
	event_ent "calend/internal/modules/db/ent/event"
	"calend/internal/modules/db/ent/invitation"
	"calend/internal/modules/db/ent/predicate"
	"calend/internal/modules/db/ent/tag"
	"calend/internal/modules/db/ent/user"
	ev_dto "calend/internal/modules/domain/event/dto"
	"calend/internal/modules/domain/event/repo"
	"calend/internal/modules/domain/search/dto"
	search_pkg "calend/internal/pkg/search/engine/db"
	"calend/internal/pkg/search/engine/db/types"
	"calend/internal/pkg/search/paginate"
	"context"
	"strings"
)

func (r *Evaluator) searchEventsDB(ctx context.Context, searchRequest *dto.EventSearchRequest) (ev_dto.Events, error) {
	filterPredicates := r.buildEventFiltersDB(searchRequest.Filter)
	sortPredicates := r.buildEventSortsDB(searchRequest.Sort)
	limit, offset := paginate.BuildPaginate(searchRequest.Paginate)

	eventsQuery := r.dbClient.Event.Query().
		Where(filterPredicates...).
		Order(sortPredicates...).
		Limit(limit).
		Offset(offset)

	events, err := eventsQuery.All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	eventUuids := make([]string, len(events))
	for i, event := range events {
		eventUuids[i] = event.ID
	}

	creatorUuids := make([]string, len(events))
	for i, event := range events {
		creatorUuids[i] = event.CreatorUUID
	}

	_, err = r.dbClient.Invitation.Query().Where(invitation.EventUUIDIn(eventUuids...)).All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	_, err = r.dbClient.Tag.Query().Where(tag.HasEventsWith(event_ent.IDIn(eventUuids...))).All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	_, err = r.dbClient.User.Query().Where(user.IDIn(creatorUuids...)).All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return repo.ToEventDTOs(events), nil
}

func (r *Evaluator) buildEventFiltersDB(f *dto.EventFilter) []predicate.Event {
	if f == nil {
		return nil
	}

	ftsSearch := []string{
		"name",
		"description",
		"type",
		"creator.login",
		"tags.name",
	}

	// Составляем карту для получения полей из связанных сущностей
	wrappersMap := map[string]types.Wrapper{
		"creator.login": func(p types.Predicate) types.Predicate {
			return types.Predicate(event_ent.HasCreatorWith(predicate.User(p)))
		},
		"tags.name": func(p types.Predicate) types.Predicate {
			return types.Predicate(event_ent.HasTagsWith(predicate.Tag(p)))
		},
		"invitations.user_uuid": func(p types.Predicate) types.Predicate {
			return types.Predicate(event_ent.HasInvitationsWith(predicate.Invitation(p)))
		},
	}

	builder := search_pkg.NewQueryBuilder(wrappersMap)

	builder.AddField("timestamp", f.Timestamp)
	builder.AddField("name", f.Name)
	builder.AddField("description", f.Description)
	builder.AddField("type", f.Type)
	builder.AddField("is_whole_day", f.IsWholeDay)
	builder.AddField("creator_uuid", f.CreatorUuid)
	builder.AddField("creator.login", f.CreatorLogin)
	builder.AddField("tags.name", f.TagName)
	builder.AddField("invitations.user_uuid", f.InvitedUserUuid)
	builder.AddField(strings.Join(ftsSearch, " "), f.FTSearchStr)

	predicates := builder.Build()

	var eventPreds []predicate.Event
	for _, p := range predicates {
		eventPreds = append(eventPreds, predicate.Event(p))
	}

	return eventPreds
}

func (r *Evaluator) buildEventSortsDB(f *dto.EventSort) []event_ent.Order {
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
