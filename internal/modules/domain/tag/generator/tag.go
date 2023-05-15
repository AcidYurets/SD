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

// Generate создает теги со случайными данными
func (r *TagGenerator) Generate(ctx context.Context, count uint) error {
	for i := uint(0); i < count; i++ {
		newTag, err := r.generateTag(i)
		if err != nil {
			return fmt.Errorf("ошибка при генерации тега: %w", err)
		}

		_, err = r.tagRepo.Create(ctx, newTag)
		if err != nil {
			return fmt.Errorf("ошибка при создании тега: %w", err)
		}
	}
	return nil
}

func (r *TagGenerator) generateTag(num uint) (*dto.CreateTag, error) {
	// Создаем тег со случайным названием и описанием
	newTag := &dto.CreateTag{
		Name:        fmt.Sprintf("Тег%d", num),
		Description: fmt.Sprintf("Описание тега %d", num),
	}

	return newTag, nil
}
