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
	"os"
	"time"
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

	file, err := os.Create("evaluations2.txt")
	if err != nil {
		fmt.Printf("ошибка создания файла: %s", err)
	}
	defer file.Close()

	fmt.Println("================== Запускаем замеры поиска ==================")
	pageSizes := []int{5, 10, 20, 50, 100, 200, 400, 500, 1000, 2000, 5000}
	count := 100

	_, err = file.Write([]byte(fmt.Sprintf("%5s%15s%15s\n", "Count", "DB", "Elastic")))
	if err != nil {
		fmt.Printf("ошибка записи в файл: %s", err)
	}

	for _, pageSize := range pageSizes {
		var sumDB time.Duration
		var sumElastic time.Duration
		for i := 0; i < count; i++ {
			evaluationRequest := &evaluator.EvaluationRequest{PageSize: pageSize}
			res, err := eval.EvaluateSearchEvents(ctx, evaluationRequest)
			if err != nil {
				fmt.Printf("ошибка выполнения замеров поиска: %s", err)
			}

			sumDB += res.DurationDB
			sumElastic += res.DurationElastic
		}

		resDB := (sumDB / time.Duration(count)).Microseconds()
		resElastic := (sumElastic / time.Duration(count)).Microseconds()
		_, err = file.Write([]byte(fmt.Sprintf("%5d%15d%15d\n", pageSize, resDB, resElastic)))
		if err != nil {
			fmt.Printf("ошибка записи в файл: %s", err)
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
