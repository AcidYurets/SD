package dataloader

import (
	"calend/internal/models/access"
	"calend/internal/modules/domain/access_right/dto"
	ar_serv "calend/internal/modules/domain/access_right/service"
	auth_serv "calend/internal/modules/domain/auth/service"
	event_serv "calend/internal/modules/domain/event/service"
	tag_serv "calend/internal/modules/domain/tag/service"
	user_dto "calend/internal/modules/domain/user/dto"
	user_serv "calend/internal/modules/domain/user/service"
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
)

type ServerLoader struct {
	userService  *user_serv.UserService
	tagService   *tag_serv.TagService
	authService  *auth_serv.AuthService
	arService    *ar_serv.AccessRightService
	eventService *event_serv.EventService
}

func (l *ServerLoader) GetAccessRightsByCodes(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	codes := KeysToAccessRightSlice(keys)

	all, err := l.arService.ListByCodes(ctx, codes)
	if err != nil {
		for index := range keys {
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
		return output
	}

	allByCode := map[access.Type]*dto.AccessRight{}

	for _, dtm := range all {
		allByCode[dtm.Code] = dtm
	}

	for index, key := range keys {
		val, ok := allByCode[access.Type(key.String())]
		if ok {
			output[index] = &dataloader.Result{Data: val, Error: nil}
		} else {
			err := fmt.Errorf("право доступа с кодом = %s не найден", key.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

func (l *ServerLoader) GetUsersByUuids(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))
	uuids := KeysToStringSlice(keys)

	all, err := l.userService.ListByUuids(ctx, uuids)
	if err != nil {
		for index := range keys {
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
		return output
	}

	allByUuid := map[string]*user_dto.User{}

	for _, dtm := range all {
		allByUuid[dtm.Uuid] = dtm
	}

	for index, key := range keys {
		val, ok := allByUuid[key.String()]
		if ok {
			output[index] = &dataloader.Result{Data: val, Error: nil}
		} else {
			err := fmt.Errorf("пользователь с uuid = %s не найден", key.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}
