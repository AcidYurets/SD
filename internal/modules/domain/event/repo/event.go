package repo

import (
	"calend/internal/modules/db"
	"calend/internal/modules/db/ent"
	event_ent "calend/internal/modules/db/ent/event"
	inv_ent "calend/internal/modules/db/ent/invitation"
	user_ent "calend/internal/modules/db/ent/user"
	"calend/internal/modules/domain/event/dto"
	tag_dto "calend/internal/modules/domain/tag/dto"
	"calend/internal/modules/domain/tag/repo"
	repo2 "calend/internal/modules/domain/user/repo"
	"context"
)

type EventRepo struct {
	client *ent.Client
}

func NewEventRepo(client *ent.Client) *EventRepo {
	return &EventRepo{
		client: client,
	}
}

func (r *EventRepo) GetByUuid(ctx context.Context, uuid string) (*dto.Event, error) {
	event, err := r.client.Event.Query().Where(event_ent.ID(uuid)).
		WithInvitations().
		Order(event_ent.ByDescription()).
		Only(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToEventDTO(event), nil
}

func (r *EventRepo) ListTagsByEventUuid(ctx context.Context, uuid string) (tag_dto.Tags, error) {
	tags, err := r.client.Event.Query().Where(event_ent.ID(uuid)).
		QueryTags().
		All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return repo.ToTagDTOs(tags), nil
}

func (r *EventRepo) GetCheckingInfoByUuid(ctx context.Context, uuid string) (*dto.Event, error) {
	event, err := r.client.Event.Query().Where(event_ent.ID(uuid)).
		WithInvitations(func(q *ent.InvitationQuery) {
			q.Select("user_uuid", "access_right_code", "event_uuid")
		}).
		Select("creator_uuid").
		Only(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToEventDTO(event), nil
}

func (r *EventRepo) ListAvailable(ctx context.Context, userUuid string) (dto.Events, error) {
	events, err := r.client.Event.Query().
		Where(
			event_ent.Or(
				// Либо переданный пользователь - создатель
				event_ent.HasCreatorWith(
					user_ent.ID(userUuid),
				),
				// Либо существует приглашение с этим пользователем
				event_ent.HasInvitationsWith(
					inv_ent.HasUserWith(user_ent.ID(userUuid)),
				),
			),
		).
		WithInvitations().
		All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToEventDTOs(events), nil
}

func (r *EventRepo) Create(ctx context.Context, dtm *dto.CreateEvent) (*dto.Event, error) {
	event, err := r.client.Event.Create().
		SetTimestamp(dtm.Timestamp).
		SetName(dtm.Name).
		SetNillableDescription(dtm.Description).
		SetType(dtm.Type).
		SetIsWholeDay(dtm.IsWholeDay).
		AddTagIDs(dtm.TagUuids...).
		SetCreatorID(dtm.CreatorUuid).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToEventDTO(event), nil
}

func (r *EventRepo) Update(ctx context.Context, uuid string, dtm *dto.UpdateEvent) (*dto.Event, error) {
	event, err := r.client.Event.UpdateOneID(uuid).
		SetTimestamp(dtm.Timestamp).
		SetName(dtm.Name).
		SetNillableDescription(dtm.Description).
		SetType(dtm.Type).
		SetIsWholeDay(dtm.IsWholeDay).
		ClearTags().
		AddTagIDs(dtm.TagUuids...).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToEventDTO(event), nil
}

func (r *EventRepo) Delete(ctx context.Context, uuid string) error {
	err := r.client.Event.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return db.WrapError(err)
	}

	return nil
}

func (r *EventRepo) SubSet(ctx context.Context, offset, limit int, uuids ...string) (dto.Events, error) {
	eventsQuery := r.client.Event.Query()

	// Если переданы uuids, то ограничиваемся ими
	if len(uuids) > 0 {
		eventsQuery = eventsQuery.Where(event_ent.IDIn(uuids...))
	}

	events, err := eventsQuery.
		Order(event_ent.ByID()).
		Limit(limit).
		Offset(offset).
		WithInvitations().
		WithTags().
		WithCreator().
		All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToEventDTOs(events), nil
}

func ToEventDTO(model *ent.Event) *dto.Event {
	if model == nil {
		return nil
	}
	return &dto.Event{
		Uuid:        model.ID,
		Timestamp:   model.Timestamp,
		Name:        model.Name,
		Description: model.Description,
		Type:        model.Type,
		IsWholeDay:  model.IsWholeDay,
		Invitations: ToInvitationDTOs(model.Edges.Invitations),
		CreatorUuid: model.CreatorUUID,

		Tags:    repo.ToTagDTOs(model.Edges.Tags),
		Creator: repo2.ToUserDTO(model.Edges.Creator),
	}
}

func ToEventDTOs(models ent.Events) dto.Events {
	if models == nil {
		return nil
	}
	dtms := make(dto.Events, len(models))
	for i := range models {
		dtms[i] = ToEventDTO(models[i])
	}
	return dtms
}
