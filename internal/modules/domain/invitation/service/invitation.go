package service

import (
	"calend/internal/modules/domain/invitation/dto"
	"context"
)

//go:generate mockgen -destination mock_test.go -package service . IInvitationRepo

type IInvitationRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.Invitation, error)
	ListByEventUuid(ctx context.Context, userUuid string) (dto.Invitations, error)
	CreateBulk(ctx context.Context, dtms dto.CreateInvitations) (dto.Invitations, error)
	Update(ctx context.Context, uuid string, dtm *dto.UpdateInvitation) (*dto.Invitation, error)
	Delete(ctx context.Context, uuid string) error
}

type InvitationService struct {
	repo IInvitationRepo
}

func NewInvitationService(repo IInvitationRepo) *InvitationService {
	return &InvitationService{
		repo: repo,
	}
}

func (r *InvitationService) GetByUuid(ctx context.Context, uuid string) (*dto.Invitation, error) {
	return r.repo.GetByUuid(ctx, uuid)
}

func (r *InvitationService) ListByAvailableEventUuid(ctx context.Context, userUuid string) (dto.Invitations, error) {
	// TODO: проверка, что событие доступно пользователю

	return r.repo.ListByEventUuid(ctx, userUuid)
}

func (r *InvitationService) CreateBulk(ctx context.Context, dtms dto.CreateInvitations) (dto.Invitations, error) {
	return r.repo.CreateBulk(ctx, dtms)
}

func (r *InvitationService) Update(ctx context.Context, uuid string, dtm *dto.UpdateInvitation) (*dto.Invitation, error) {
	return r.repo.Update(ctx, uuid, dtm)
}

func (r *InvitationService) Delete(ctx context.Context, uuid string) error {
	return r.repo.Delete(ctx, uuid)
}
