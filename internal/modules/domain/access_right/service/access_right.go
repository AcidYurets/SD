package service

import (
	"calend/internal/models/access"
	"calend/internal/modules/domain/access_right/dto"
	"context"
)

//go:generate mockgen -destination mock_test.go -package service . IAccessRightRepo

type IAccessRightRepo interface {
	GetByCode(ctx context.Context, code access.Type) (*dto.AccessRight, error)
	ListByCodes(ctx context.Context, code []access.Type) (dto.AccessRights, error)
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

func (r *AccessRightService) GetByCode(ctx context.Context, code access.Type) (*dto.AccessRight, error) {
	return r.repo.GetByCode(ctx, code)
}

func (r *AccessRightService) ListByCodes(ctx context.Context, codes []access.Type) (dto.AccessRights, error) {
	return r.repo.ListByCodes(ctx, codes)
}

func (r *AccessRightService) List(ctx context.Context) (dto.AccessRights, error) {
	return r.repo.List(ctx)
}
