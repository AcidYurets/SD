package integrational

import (
	"calend/internal/modules/db/ent"
	"calend/internal/modules/domain/access_right/dto"
	"calend/internal/modules/domain/access_right/service"
	auth_dto "calend/internal/modules/domain/auth/dto"
	auth_serv "calend/internal/modules/domain/auth/service"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func accessRightServiceTest(t *testing.T, service *service.AccessRightService, authService *auth_serv.AuthService, client *ent.Client) {
	err := truncateAll(client)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	err = createAccessRight(client)
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

	existingAr := dto.AccessRights{
		&dto.AccessRight{
			Code:        "r",
			Description: "Право на просмотр",
		},
		&dto.AccessRight{
			Code:        "ri",
			Description: "Право на просмотр и приглашение",
		},
	}

	ar1, err := service.GetByCode(ctx, "ri")
	assert.NoError(t, err)
	assert.Equal(t, existingAr[1], ar1)

	ars, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, existingAr[0], ars[0])
}
