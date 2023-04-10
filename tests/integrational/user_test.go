package integrational

import (
	"calend/internal/modules/db/ent"
	auth_dto "calend/internal/modules/domain/auth/dto"
	auth_serv "calend/internal/modules/domain/auth/service"
	"calend/internal/modules/domain/user/dto"
	user_serv "calend/internal/modules/domain/user/service"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func userServiceTest(t *testing.T, userService *user_serv.UserService, authService *auth_serv.AuthService, client *ent.Client) {
	err := truncateAll(client)
	assert.NoError(t, err)
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

	users1, err := userService.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, currentUser, users1[0])

	currentUser.Login += "1"
	updateUser := &dto.UpdateUser{
		Phone: currentUser.Phone,
		Login: currentUser.Login,
	}
	u2, err := userService.Update(ctx, currentUser.Uuid, updateUser)
	assert.NoError(t, err)
	assert.Equal(t, currentUser, u2)

	err = userService.Delete(ctx, currentUser.Uuid)
	assert.NoError(t, err)

	// После удаление по Uuid все еще можно получить запись
	u3, err := userService.GetByUuid(ctx, currentUser.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, currentUser, u3)

	// А в списке ее не будет
	users2, err := userService.List(ctx)
	assert.NoError(t, err)
	assert.Empty(t, users2)

	u4, err := userService.Restore(ctx, currentUser.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, currentUser, u4)

	// Восстанавливаем и в списке она тоже есть
	users3, err := userService.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, currentUser, users3[0])
}
