package service

import (
	"calend/internal/models/roles"
	"calend/internal/modules/domain/user/dto"
	"context"
)

//go:generate mockgen -destination mock_test.go -package service . IUserRepo

type IUserRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.User, error)
	List(ctx context.Context) (dto.Users, error)
	Update(ctx context.Context, uuid string, dtm *dto.UpdateUser) (*dto.User, error)
	Delete(ctx context.Context, uuid string) error
	Restore(ctx context.Context, uuid string) (*dto.User, error)
}

type UserService struct {
	repo IUserRepo
}

func NewUserService(repo IUserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (r *UserService) GetByUuid(ctx context.Context, uuid string) (*dto.User, error) {
	return r.repo.GetByUuid(ctx, uuid)
}

func (r *UserService) List(ctx context.Context) (dto.Users, error) {
	return r.repo.List(ctx)
}

func (r *UserService) Update(ctx context.Context, uuid string, dtm *dto.UpdateUser) (*dto.User, error) {
	return r.repo.Update(ctx, uuid, dtm)
}

func (r *UserService) ChangeRole(ctx context.Context, uuid string, role roles.Type) (*dto.User, error) {
	dtm := &dto.UpdateUser{
		Role: &role,
	}

	return r.repo.Update(ctx, uuid, dtm)
}

func (r *UserService) Delete(ctx context.Context, uuid string) error {
	return r.repo.Delete(ctx, uuid)
}

func (r *UserService) Restore(ctx context.Context, uuid string) (*dto.User, error) {
	return r.repo.Restore(ctx, uuid)
}
