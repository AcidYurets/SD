package integrational

import (
	"calend/internal/modules/db/ent"
	auth_dto "calend/internal/modules/domain/auth/dto"
	auth_serv "calend/internal/modules/domain/auth/service"
	"calend/internal/modules/domain/event/dto"
	event_serv "calend/internal/modules/domain/event/service"
	search_dto "calend/internal/modules/domain/search/dto"
	"calend/internal/modules/domain/search/service"
	tag_dto "calend/internal/modules/domain/tag/dto"
	tag_serv "calend/internal/modules/domain/tag/service"
	"calend/internal/pkg/search/filter"
	"calend/internal/pkg/search/paginate"
	"calend/internal/pkg/search/sort"
	"calend/internal/utils/ptr"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func searchServiceTest(t *testing.T, service *service.SearchService, eventService *event_serv.EventService, tagService *tag_serv.TagService, authService *auth_serv.AuthService, client *ent.Client) {
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
	_ = makeCtxByUser(currentUser2)
	_ = makeCtxByUser(currentUser3)

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
	_, err = tagService.Create(ctx1, newTag2)

	// Пользователь 1 создает событие и приглашает пользователя 2 с правами на чтение
	newEvent := &dto.CreateEvent{
		Timestamp:   time.Now().Add(time.Hour),
		Name:        "Конец лекции",
		Description: ptr.String("Наконец-то я смогу пойти домой"),
		Type:        "Обычное",
		IsWholeDay:  false,
		TagUuids:    []string{tag1.Uuid},
	}
	newInvs := dto.CreateEventInvitations{{
		UserUuid:        currentUser2.Uuid,
		AccessRightCode: "r",
	}}
	_, err = eventService.CreateWithInvitations(ctx1, newEvent, newInvs)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	// Производим поиск события
	sortDir := sort.DirectionDesc
	searchRequest := &search_dto.EventSearchRequest{
		Filter: &search_dto.EventFilter{
			Name: &filter.TextQueryFilter{
				Ts: ptr.String("Кон"),
			},
			Description: &filter.TextQueryFilter{
				Ts: ptr.String("смо"),
			},
		},
		Sort: &search_dto.EventSort{
			Timestamp:   &sortDir,
			Description: &sortDir,
		},
		Paginate: &paginate.ByPage{
			Page:     ptr.Int(1),
			PageSize: ptr.Int(20),
		},
	}
	es1, err := service.SearchEvents(ctx1, searchRequest)
	assert.NoError(t, err)
	assert.Len(t, es1, 1)
	fmt.Println(es1)

}
