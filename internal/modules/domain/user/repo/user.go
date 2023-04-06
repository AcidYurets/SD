package service

import (
	"calend/internal/modules/db"
	"calend/internal/modules/db/ent"
	"calend/internal/modules/domain/user/dto"
	"context"
	"time"
)

type UserRepo struct {
	client *ent.Client
}

func NewUserRepo(client *ent.Client) *UserRepo {
	return &UserRepo{
		client: client,
	}
}

func (r *UserRepo) GetByUuid(ctx context.Context, uuid string) (*dto.User, error) {
	user, err := r.client.User.Get(ctx, uuid)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return toDTO(user), nil
}

func (r *UserRepo) List(ctx context.Context) (dto.Users, error) {
	users, err := r.client.User.Query().All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return toDTOs(users), nil
}

func (r *UserRepo) Update(ctx context.Context, uuid string, dtm *dto.UpdateUser) (*dto.User, error) {
	user, err := r.client.User.UpdateOneID(uuid).
		SetPhone(dtm.Phone).
		SetLogin(dtm.Login).
		SetPasswordHash(dtm.PasswordHash).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return toDTO(user), nil
}

func (r *UserRepo) Delete(ctx context.Context, uuid string) error {
	err := r.client.User.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return db.WrapError(err)
	}

	return nil
}

func (r *UserRepo) Restore(ctx context.Context, uuid string) (*dto.User, error) {
	user, err := r.client.User.UpdateOneID(uuid).SetDeletedAt(time.Time{}).Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return toDTO(user), nil
}

func toDTO(model *ent.User) *dto.User {
	return &dto.User{
		Uuid:         model.ID,
		Phone:        model.Phone,
		Login:        model.Login,
		PasswordHash: model.PasswordHash,
	}
}

func toDTOs(models ent.Users) dto.Users {
	if models == nil {
		return nil
	}
	dtms := make(dto.Users, len(models))
	for i := range models {
		dtms[i] = toDTO(models[i])
	}
	return dtms
}
