package integrational

import (
	"calend/internal/modules/db/ent"
	auth_dto "calend/internal/modules/domain/auth/dto"
	auth_serv "calend/internal/modules/domain/auth/service"
	"calend/internal/modules/domain/event/dto"
	event_serv "calend/internal/modules/domain/event/service"
	tag_dto "calend/internal/modules/domain/tag/dto"
	tag_serv "calend/internal/modules/domain/tag/service"
	"calend/internal/utils/ptr"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func eventServiceTest(t *testing.T, service *event_serv.EventService, tagService *tag_serv.TagService, authService *auth_serv.AuthService, client *ent.Client) {
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

	// Регистрируем пользователей
	newUser1 := &auth_dto.NewUser{
		Login:    "yurets",
		Password: "Pass123!",
		Phone:    "+79197628803",
	}
	currentUser1, err := authService.SignUp(context.Background(), newUser1)
	assert.NoError(t, err)
	newUser2 := &auth_dto.NewUser{
		Login:    "nekit",
		Password: "Pass1234!",
		Phone:    "+79197621234",
	}
	currentUser2, err := authService.SignUp(context.Background(), newUser2)
	assert.NoError(t, err)
	newUser3 := &auth_dto.NewUser{
		Login:    "sasha",
		Password: "Pass12345!",
		Phone:    "+79197624321",
	}
	currentUser3, err := authService.SignUp(context.Background(), newUser3)
	assert.NoError(t, err)
	// Возврат, т.к. не получится создать сессию с пользователем
	if err != nil {
		return
	}

	// Создаем контексты с сессиями
	ctx1 := makeCtxByUser(currentUser1)
	ctx2 := makeCtxByUser(currentUser2)
	ctx3 := makeCtxByUser(currentUser3)

	// Создадим теги
	newTag1 := &tag_dto.CreateTag{
		Name:        "Для друзей",
		Description: "Что-то только для друзей",
	}
	newTag2 := &tag_dto.CreateTag{
		Name:        "Для себя",
		Description: "Что-то личное",
	}
	tag1, err := tagService.Create(ctx1, newTag1)
	tag2, err := tagService.Create(ctx1, newTag2)

	// Пользователь 1 создает событие и приглашает пользователя 2 с правами на чтение
	newEvent := &dto.CreateEvent{
		Timestamp:   time.Now().Add(time.Hour),
		Name:        "Конец лекции",
		Description: ptr.String("Наконец-то я смогу пойти домой"),
		Type:        "Обычное",
		IsWholeDay:  false,
		TagUuids:    []string{tag1.Uuid},
	}
	newInvs := dto.CreateInvitations{{
		UserUuid:        currentUser2.Uuid,
		AccessRightCode: "r",
	}}
	event, err := service.Create(ctx1, newEvent, newInvs)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	// Получаем событие 1ым пользователем
	e1, err := service.GetByUuid(ctx1, event.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, event, e1)

	es1, err := service.ListAvailable(ctx1)
	assert.NoError(t, err)
	assert.Equal(t, event, es1[0])

	// Получаем событие 2ым пользователем
	e2, err := service.GetByUuid(ctx2, event.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, event, e2)

	es2, err := service.ListAvailable(ctx2)
	assert.NoError(t, err)
	assert.Equal(t, event, es2[0])

	// Меняем имя и тег, добавляем приглашение
	updEvent := &dto.UpdateEvent{
		Timestamp:   event.Timestamp,
		Name:        event.Name + "1",
		Description: event.Description,
		Type:        event.Type,
		IsWholeDay:  event.IsWholeDay,
		TagUuids:    []string{tag2.Uuid},
	}
	updInvs := dto.CreateInvitations{
		&dto.CreateInvitation{
			UserUuid:        currentUser2.Uuid,
			AccessRightCode: "r",
		},
		&dto.CreateInvitation{
			UserUuid:        currentUser3.Uuid,
			AccessRightCode: "riu",
		},
	}

	// Изменяем событие 1ым пользователем
	event, err = service.Update(ctx1, event.Uuid, updEvent, updInvs)
	assert.NoError(t, err)

	// Проверяем корректность нового тега
	tags, err := service.ListTagsByEventUuid(ctx1, event.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, tags[0], tag2)

	// Пробуем изменить событие 2ым пользователем
	_, err = service.Update(ctx2, event.Uuid, updEvent, updInvs)
	assert.Error(t, err)

	// Изменяем событие 3им пользователем
	_, err = service.Update(ctx3, event.Uuid, updEvent, updInvs)
	assert.NoError(t, err)

	// Пробуем удалить событие 3им пользователем
	err = service.Delete(ctx3, event.Uuid)
	assert.Error(t, err)

	// Удаляем событие 1ым пользователем
	err = service.Delete(ctx1, event.Uuid)
	assert.NoError(t, err)

	// Заново создаем событие с 1 приглашением
	event, err = service.Create(ctx1, newEvent, newInvs)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	addInvs1 := dto.CreateInvitations{
		&dto.CreateInvitation{
			UserUuid:        currentUser3.Uuid,
			AccessRightCode: "ri",
		},
	}

	addInvs2 := dto.CreateInvitations{
		&dto.CreateInvitation{
			UserUuid:        currentUser3.Uuid,
			AccessRightCode: "riud",
		},
	}

	// Пробуем добавить приглашение 2ым пользователем
	_, err = service.AddInvitations(ctx2, event.Uuid, addInvs1)
	assert.Error(t, err)

	// Добавляем приглашение 1ым пользователем
	_, err = service.AddInvitations(ctx1, event.Uuid, addInvs1)
	assert.NoError(t, err)

	// Пробуем добавить уже существующее приглашение 3им пользователем
	_, err = service.AddInvitations(ctx3, event.Uuid, addInvs2)
	assert.Error(t, err)
}
