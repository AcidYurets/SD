package service

import (
	"calend/internal/modules/domain/tag/dto"
	"context"
)

//go:generate mockgen -destination mock_test.go -package service . ITagRepo

type ITagRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.Tag, error)
	List(ctx context.Context) (dto.Tags, error)
	Create(ctx context.Context, dtm *dto.CreateTag) (*dto.Tag, error)
	Update(ctx context.Context, uuid string, dtm *dto.UpdateTag) (*dto.Tag, error)
	Delete(ctx context.Context, uuid string) error
	Restore(ctx context.Context, uuid string) (*dto.Tag, error)
}

type TagService struct {
	repo ITagRepo
}

func NewTagService(repo ITagRepo) *TagService {
	return &TagService{
		repo: repo,
	}
}

func (r *TagService) GetByUuid(ctx context.Context, uuid string) (*dto.Tag, error) {
	return r.repo.GetByUuid(ctx, uuid)
}

func (r *TagService) List(ctx context.Context) (dto.Tags, error) {
	return r.repo.List(ctx)
}

func (r *TagService) Create(ctx context.Context, dtm *dto.CreateTag) (*dto.Tag, error) {
	return r.repo.Create(ctx, dtm)
}

func (r *TagService) Update(ctx context.Context, uuid string, dtm *dto.UpdateTag) (*dto.Tag, error) {
	return r.repo.Update(ctx, uuid, dtm)
}

func (r *TagService) Delete(ctx context.Context, uuid string) error {
	return r.repo.Delete(ctx, uuid)
}

func (r *TagService) Restore(ctx context.Context, uuid string) (*dto.Tag, error) {
	return r.repo.Restore(ctx, uuid)
}
