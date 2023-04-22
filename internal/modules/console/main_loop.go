package console

import (
	"calend/internal/models/session"
	"calend/internal/modules/db"
	"calend/internal/modules/db/ent"
	"calend/internal/modules/db/seed"
	ar_serv "calend/internal/modules/domain/access_right/service"
	auth_dto "calend/internal/modules/domain/auth/dto"
	auth_serv "calend/internal/modules/domain/auth/service"
	dto2 "calend/internal/modules/domain/event/dto"
	event_serv "calend/internal/modules/domain/event/service"
	search_serv "calend/internal/modules/domain/search/service"
	tag_dto "calend/internal/modules/domain/tag/dto"
	tag_serv "calend/internal/modules/domain/tag/service"
	"calend/internal/modules/domain/user/dto"
	user_serv "calend/internal/modules/domain/user/service"
	"context"
	"fmt"
	"github.com/google/uuid"
)

const menu = `
===== Многопользовательский календарь =====
0. Выйти
1. Создать событие
2. Обновить событие
3. Добавить приглашения
4. Удалить событие
5. Получить событие
6. Получить все доступные события
7. Создать тег
8. Обновить тег
9. Удалить тег
10. Получить все теги

100. Возврат БД к исходному состоянию
Что будем делать: `

func MenuLoop(
	userService *user_serv.UserService,
	tagService *tag_serv.TagService,
	authService *auth_serv.AuthService,
	arService *ar_serv.AccessRightService,
	eventService *event_serv.EventService,
	searchService *search_serv.SearchService,

	client *ent.Client,
) error {
	ctx, err := beforeStart(client, authService, tagService)
	if err != nil {
		return err
	}

	for {
		fmt.Print(menu)

		var c int
		_, err := fmt.Scan(&c)
		if err != nil {
			printError(err)
			continue
		}

		switch c {
		case 1:
			createEvent := inputCreateEvent()
			var event *dto2.Event
			event, err = eventService.Create(ctx, createEvent, nil)
			printEvents(event)

		case 2:
			id := inputUuid()
			updateEvent := inputUpdateEvent()
			var event *dto2.Event
			event, err = eventService.Update(ctx, id, updateEvent, nil)
			printEvents(event)

		case 3:
			id := inputUuid()
			invitations := inputInvitations()
			var event *dto2.Event
			event, err = eventService.AddInvitations(ctx, id, invitations)
			printEvents(event)

		case 4:
			id := inputUuid()
			err = eventService.Delete(ctx, id)

		case 5:
			id := inputUuid()
			var event *dto2.Event
			event, err = eventService.GetByUuid(ctx, id)
			printEvents(event)

		case 6:
			var events dto2.Events
			events, err = eventService.ListAvailable(ctx)
			printEvents(events...)

		case 7:
			newTag := inputCreateTag()
			var tag *tag_dto.Tag
			tag, err = tagService.Create(ctx, newTag)
			printTags(tag)

		case 8:
			id := inputUuid()
			updateTag := inputUpdateTag()
			var tag *tag_dto.Tag
			tag, err = tagService.Update(ctx, id, updateTag)
			printTags(tag)

		case 9:
			id := inputUuid()
			err = tagService.Delete(ctx, id)

		case 10:
			var tags tag_dto.Tags
			tags, err = tagService.List(ctx)
			printTags(tags...)

		case 100:
			c, err := beforeStart(client, authService, tagService)
			if err == nil {
				ctx = c
			}

		case 0:
			return nil

		}
		if err != nil {
			printError(err)
		}
	}
}

func beforeStart(client *ent.Client, authService *auth_serv.AuthService, tagService *tag_serv.TagService) (context.Context, error) {
	err := db.TruncateAll(client)
	if err != nil {
		return nil, err
	}

	err = seed.AccessRight(client)
	if err != nil {
		return nil, err
	}

	// Регистрируем пользователей
	newUser1 := &auth_dto.NewUser{
		Login:    "yurets",
		Password: "Pass123!",
		Phone:    "+79197628803",
	}
	currentUser, err := authService.SignUp(context.Background(), newUser1)
	if err != nil {
		return nil, err
	}
	ctx := makeCtxByUser(currentUser)

	// Создадим теги
	newTag1 := &tag_dto.CreateTag{
		Name:        "Для друзей",
		Description: "Что-то только для друзей",
	}
	newTag2 := &tag_dto.CreateTag{
		Name:        "Для себя",
		Description: "Что-то личное",
	}

	_, err = tagService.Create(ctx, newTag1)
	if err != nil {
		return nil, err
	}
	_, err = tagService.Create(ctx, newTag2)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}

func makeCtxByUser(user *dto.User) context.Context {
	ss := session.Session{
		SID:      uuid.NewString(),
		UserUuid: user.Uuid,
	}

	ctx := context.Background()
	return session.SetSessionToCtx(ctx, ss)
}
