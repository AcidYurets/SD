package integrational

import (
	"calend/internal/modules/db/ent"
	auth_dto "calend/internal/modules/domain/auth/dto"
	auth_serv "calend/internal/modules/domain/auth/service"
	"calend/internal/modules/domain/tag/dto"
	"calend/internal/modules/domain/tag/service"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func tagServiceTest(t *testing.T, service *service.TagService, authService *auth_serv.AuthService, client *ent.Client) {
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

	// Создаем тег
	newTag := &dto.CreateTag{
		Name:        "Праздник",
		Description: "Тег для обозначения праздников",
	}
	tag, err := service.Create(ctx, newTag)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	t1, err := service.GetByUuid(ctx, tag.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, tag, t1)

	ts1, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, tag, ts1[0])

	tag.Name += "1"
	updateUser := &dto.UpdateTag{
		Name:        tag.Name,
		Description: tag.Description,
	}
	t2, err := service.Update(ctx, tag.Uuid, updateUser)
	assert.NoError(t, err)
	assert.Equal(t, tag, t2)

	err = service.Delete(ctx, tag.Uuid)
	assert.NoError(t, err)

	// После удаление по Uuid все еще можно получить запись
	t3, err := service.GetByUuid(ctx, tag.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, tag, t3)

	// А в списке ее не будет
	ts2, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts2)

	t4, err := service.Restore(ctx, tag.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, tag, t4)

	// Восстанавливаем и в списке она тоже есть
	ts3, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, tag, ts3[0])
}
