package repo

import (
	"calend/internal/modules/db"
	"calend/internal/modules/db/ent"
	user_ent "calend/internal/modules/db/ent/user"
	"calend/internal/modules/db/schema"
	"calend/internal/modules/domain/user/dto"
	"context"
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
	user, err := r.client.User.Get(schema.SkipSoftDelete(ctx), uuid)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToUserDTO(user), nil
}

func (r *UserRepo) GetByLogin(ctx context.Context, login string) (*dto.User, error) {
	user, err := r.client.User.Query().Where(user_ent.Login(login)).Only(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToUserDTO(user), nil
}

func (r *UserRepo) List(ctx context.Context) (dto.Users, error) {
	users, err := r.client.User.Query().All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToUserDTOs(users), nil
}

func (r *UserRepo) Create(ctx context.Context, dtm *dto.CreateUser) (*dto.User, error) {
	user, err := r.client.User.Create().
		SetPhone(dtm.Phone).
		SetLogin(dtm.Login).
		SetPasswordHash(dtm.PasswordHash).
		SetRole(dtm.Role).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToUserDTO(user), nil
}

func (r *UserRepo) Update(ctx context.Context, uuid string, dtm *dto.UpdateUser) (*dto.User, error) {
	// TODO: Вместо этой фигни написать template для опциональной установки значения

	updateQuery := r.client.User.UpdateOneID(uuid)

	if dtm.Phone != nil {
		updateQuery.SetPhone(*dtm.Phone)
	}
	if dtm.Login != nil {
		updateQuery.SetLogin(*dtm.Login)
	}
	if dtm.Role != nil {
		updateQuery.SetRole(*dtm.Role)
	}

	user, err := updateQuery.Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToUserDTO(user), nil
}

func (r *UserRepo) Delete(ctx context.Context, uuid string) error {
	err := r.client.User.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return db.WrapError(err)
	}

	return nil
}

func (r *UserRepo) Restore(ctx context.Context, uuid string) (*dto.User, error) {
	user, err := r.client.User.UpdateOneID(uuid).ClearDeletedAt().Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToUserDTO(user), nil
}

func ToUserDTO(model *ent.User) *dto.User {
	if model == nil {
		return nil
	}
	return &dto.User{
		Uuid:         model.ID,
		Phone:        model.Phone,
		Login:        model.Login,
		PasswordHash: model.PasswordHash,
		Role:         model.Role,
	}
}

func ToUserDTOs(models ent.Users) dto.Users {
	if models == nil {
		return nil
	}
	dtms := make(dto.Users, len(models))
	for i := range models {
		dtms[i] = ToUserDTO(models[i])
	}
	return dtms
}
