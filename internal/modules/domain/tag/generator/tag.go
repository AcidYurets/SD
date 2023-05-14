package generator

import (
	"calend/internal/modules/domain/tag/dto"
	"context"
	"fmt"
)

type ITagRepo interface {
	// Create создает событие без приглашений, возвращает событие без связанных сущностей
	Create(ctx context.Context, dtm *dto.CreateTag) (*dto.Tag, error)
}

type TagGenerator struct {
	tagRepo ITagRepo
}

func NewTagGenerator(eRepo ITagRepo) *TagGenerator {
	return &TagGenerator{
		tagRepo: eRepo,
	}
}

func (r *TagGenerator) Generate(ctx context.Context) error {
	newTag, err := r.generateTag()
	if err != nil {
		return err
	}

	_, err = r.tagRepo.Create(ctx, newTag)
	if err != nil {
		return fmt.Errorf("ошибка при создании события: %w", err)
	}

	return nil
}

func (r *TagGenerator) generateTag() (*dto.CreateTag, error) {
	return nil, nil
}
