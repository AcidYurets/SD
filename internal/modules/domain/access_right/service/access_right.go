package service

import (
	"calend/internal/modules/domain/access_right/dto"
	"context"
)

//go:generate mockgen -destination mock_test.go -package service . IAccessRightRepo

type IAccessRightRepo interface {
	GetByCode(ctx context.Context, uuid string) (*dto.AccessRight, error)
	List(ctx context.Context) (dto.AccessRights, error)
}

type AccessRightService struct {
	repo IAccessRightRepo
}

func NewAccessRightService(repo IAccessRightRepo) *AccessRightService {
	return &AccessRightService{
		repo: repo,
	}
}

func (r *AccessRightService) GetByCode(ctx context.Context, uuid string) (*dto.AccessRight, error) {
	return r.repo.GetByCode(ctx, uuid)
}

func (r *AccessRightService) List(ctx context.Context) (dto.AccessRights, error) {
	return r.repo.List(ctx)
}
