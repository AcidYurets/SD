package generator

import (
	"calend/internal/models/roles"
	"calend/internal/modules/domain/user/dto"
	"calend/internal/utils/random"
	"context"
	"fmt"
)

type IUserRepo interface {
	// Create создает событие без приглашений, возвращает событие без связанных сущностей
	Create(ctx context.Context, dtm *dto.CreateUser) (*dto.User, error)
}

type UserGenerator struct {
	userRepo IUserRepo
}

func NewUserGenerator(eRepo IUserRepo) *UserGenerator {
	return &UserGenerator{
		userRepo: eRepo,
	}
}

// Generate создает пользователей со случайными данными
func (r *UserGenerator) Generate(ctx context.Context, count uint) error {
	for i := uint(0); i < count; i++ {
		newUser, err := r.generateUser(i)
		if err != nil {
			return fmt.Errorf("ошибка при генерации тега: %w", err)
		}

		_, err = r.userRepo.Create(ctx, newUser)
		if err != nil {
			return fmt.Errorf("ошибка при создании тега: %w", err)
		}
	}
	return nil
}

func (r *UserGenerator) generateUser(num uint) (*dto.CreateUser, error) {
	rolse := []roles.Type{"simple_user", "premium_user", "admin"}

	// Создаем пользователя
	newUser := &dto.CreateUser{
		Phone:        fmt.Sprintf("phone_%d", num),
		Login:        fmt.Sprintf("User%d", num),
		PasswordHash: "$2a$14$BOX93qyb.Pr66j4xfZRYdu8Ls98QXeKBo4zQGA.D/6QDKl6DDtEfi",
		Role:         random.FromSlice(rolse),
	}

	return newUser, nil
}
