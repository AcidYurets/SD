package repo

import (
	"calend/internal/models/access"
	"calend/internal/modules/db"
	"calend/internal/modules/db/ent"
	"calend/internal/modules/domain/access_right/dto"
	"context"
)

type AccessRightRepo struct {
	client *ent.Client
}

func NewAccessRightRepo(client *ent.Client) *AccessRightRepo {
	return &AccessRightRepo{
		client: client,
	}
}

func (r *AccessRightRepo) GetByCode(ctx context.Context, code access.Type) (*dto.AccessRight, error) {
	ar, err := r.client.AccessRight.Get(ctx, code)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToAccessRightDTO(ar), nil
}

func (r *AccessRightRepo) List(ctx context.Context) (dto.AccessRights, error) {
	ars, err := r.client.AccessRight.Query().All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToAccessRightDTOs(ars), nil
}

func ToAccessRightDTO(model *ent.AccessRight) *dto.AccessRight {
	if model == nil {
		return nil
	}
	return &dto.AccessRight{
		Code:        model.ID,
		Description: model.Description,
	}
}

func ToAccessRightDTOs(models ent.AccessRights) dto.AccessRights {
	if models == nil {
		return nil
	}
	dtms := make(dto.AccessRights, len(models))
	for i := range models {
		dtms[i] = ToAccessRightDTO(models[i])
	}
	return dtms
}
