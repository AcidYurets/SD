package integrational

import (
	"calend/internal/modules/db/ent"
	"calend/internal/modules/db/schema"
	auth_dto "calend/internal/modules/domain/auth/dto"
	auth_serv "calend/internal/modules/domain/auth/service"
	user_serv "calend/internal/modules/domain/user/service"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func authServiceTest(t *testing.T, userService *user_serv.UserService, authService *auth_serv.AuthService, client *ent.Client) {
	_, err := client.User.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	assert.NoError(t, err)
	// Если не получилось - дальше продолжать смысла нет
	if err != nil {
		return
	}

	// Регистрируем пользователя
	newUser := &auth_dto.NewUser{
		Login:    "yurets",
		Password: "Pass123!",
		Phone:    "+79197628803",
	}
	currentUser, err := authService.SignUp(context.Background(), newUser)
	assert.NoError(t, err)
	// Возврат, т.к. не получится создать сессию с пользователем
	if err != nil {
		return
	}

	// Создаем контекст с сессией
	ctx := makeCtxByUser(currentUser)

	u1, err := userService.GetByUuid(ctx, currentUser.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, currentUser, u1)

	// Логинимся с правильным паролем
	userCorrect := &auth_dto.UserCredentials{
		Login:    "yurets",
		Password: "Pass123!",
	}
	jwt, err := authService.Login(ctx, userCorrect)
	assert.NoError(t, err)
	assert.Equal(t, currentUser.Uuid, jwt.Session.UserUuid)

	// Логинимся с неправильным паролем
	userIncorrect := &auth_dto.UserCredentials{
		Login:    "yurets",
		Password: "Pass1234!",
	}
	_, err = authService.Login(ctx, userIncorrect)
	assert.Error(t, err)
}
