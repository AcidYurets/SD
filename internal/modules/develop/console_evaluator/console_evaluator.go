package console_evaluator

import (
	"calend/internal/models/session"
	"calend/internal/modules/develop/evaluator"
	"calend/internal/modules/domain/user/dto"
	user_serv "calend/internal/modules/domain/user/service"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/fx"
)

func Eval(
	userService *user_serv.UserService,
	eval *evaluator.Evaluator,

	lifecycle fx.Lifecycle,
	shutdowner fx.Shutdowner,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				evaluate(userService, eval)

				_ = shutdowner.Shutdown()
			}()

			return nil
		},
	})
}

func evaluate(
	userService *user_serv.UserService,
	eval *evaluator.Evaluator,
) {
	// Получаем пользователя yurets
	user, err := userService.GetByUuid(context.Background(), "c635e505-f957-4e33-b065-694e9cdf4e5c")
	if err != nil {
		fmt.Printf("ошибка получения пользователя по Uuid: %s", err)
	}
	ctx := makeCtxByUser(user)

	fmt.Println("================== Запускаем замеры поиска ==================")
	pageSizes := []int{20, 50, 100, 500, 1000, 2000, 5000}
	count := 10

	for _, pageSize := range pageSizes {
		for i := 0; i < count; i++ {
			evaluationRequest := &evaluator.EvaluationRequest{PageSize: pageSize}
			res, err := eval.EvaluateSearchEvents(ctx, evaluationRequest)
			if err != nil {
				fmt.Printf("ошибка выполнения замеров поиска: %s", err)
			}

			fmt.Printf("%d page size, №%d: %s\n", pageSize, i, res)
		}
	}

	fmt.Println("================== Завершаем замеры поиска ==================")
}

func makeCtxByUser(user *dto.User) context.Context {
	ss := session.Session{
		SID:      uuid.NewString(),
		UserUuid: user.Uuid,
		Role:     user.Role,
	}

	ctx := context.Background()
	return session.SetSessionToCtx(ctx, ss)
}
