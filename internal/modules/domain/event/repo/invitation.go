package repo

import (
	"calend/internal/modules/db"
	"calend/internal/modules/db/ent"
	event_ent "calend/internal/modules/db/ent/event"
	inv_ent "calend/internal/modules/db/ent/invitation"
	ar_repo "calend/internal/modules/domain/access_right/repo"
	"calend/internal/modules/domain/event/dto"
	user_repo "calend/internal/modules/domain/user/repo"
	"context"
)

type InvitationRepo struct {
	client *ent.Client
}

func NewInvitationRepo(client *ent.Client) *InvitationRepo {
	return &InvitationRepo{
		client: client,
	}
}

func (r *InvitationRepo) CreateBulk(ctx context.Context, dtms dto.CreateInvitations) (dto.Invitations, error) {
	bulk := make([]*ent.InvitationCreate, len(dtms))
	for i, dtm := range dtms {
		bulk[i] = r.client.Invitation.Create().
			SetEventID(dtm.EventUuid).
			SetUserID(dtm.UserUuid).
			SetAccessRightID(dtm.AccessRightCode)
	}

	invitations, err := r.client.Invitation.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToInvitationDTOs(invitations), nil
}

func (r *InvitationRepo) DeleteByEventUuid(ctx context.Context, eventUuid string) (int, error) {
	n, err := r.client.Invitation.Delete().Where(
		inv_ent.HasEventWith(event_ent.ID(eventUuid)),
	).Exec(ctx)
	if err != nil {
		return 0, db.WrapError(err)
	}

	return n, nil
}

func ToInvitationDTO(model *ent.Invitation) *dto.Invitation {
	if model == nil {
		return nil
	}
	return &dto.Invitation{
		Uuid:        model.ID,
		User:        user_repo.ToUserDTO(model.Edges.User),
		AccessRight: ar_repo.ToAccessRightDTO(model.Edges.AccessRight),
	}
}

func ToInvitationDTOs(models ent.Invitations) dto.Invitations {
	if models == nil {
		return nil
	}
	dtms := make(dto.Invitations, len(models))
	for i := range models {
		dtms[i] = ToInvitationDTO(models[i])
	}
	return dtms
}
