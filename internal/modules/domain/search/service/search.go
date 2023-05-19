package service

import (
	"calend/internal/models/roles"
	"calend/internal/models/session"
	ev_dto "calend/internal/modules/domain/event/dto"
	"calend/internal/modules/domain/search/dto"
	"calend/internal/pkg/search/filter"
	"context"
	"fmt"
)

//go:generate mockgen -destination mock_test.go -package service . ISearchRepo

type ISearchRepo interface {
	SearchEvents(ctx context.Context, searchRequest *dto.EventSearchRequest) (ev_dto.Events, error)
}

type IEventRepo interface {
	// GetCheckingInfoByUuid получение по uuid события только необходимыми для проверки прав доступа полями
	GetCheckingInfoByUuid(ctx context.Context, uuid string) (*ev_dto.Event, error)
}

type SearchService struct {
	repo      ISearchRepo
	eventRepo IEventRepo
}

func NewSearchService(repo ISearchRepo, eventRepo IEventRepo) *SearchService {
	return &SearchService{
		repo:      repo,
		eventRepo: eventRepo,
	}
}

func (r *SearchService) SearchEvents(ctx context.Context, searchRequest *dto.EventSearchRequest) (ev_dto.Events, error) {
	if err := r.addFilters(ctx, searchRequest); err != nil {
		return nil, fmt.Errorf("ошибка при добавлении фильтров: %w", err)
	}

	return r.repo.SearchEvents(ctx, searchRequest)
}

func (r *SearchService) addFilters(ctx context.Context, searchRequest *dto.EventSearchRequest) error {
	ss, ok := session.GetSessionFromCtx(ctx)
	if !ok {
		return fmt.Errorf("ошибка при получении сессии")
	}

	// Если пользователь -- админ, то у него полный доступ
	if ss.Role == roles.Admin {
		return nil
	}

	userUuid := ss.UserUuid

	if searchRequest == nil {
		searchRequest = &dto.EventSearchRequest{}
	}
	if searchRequest.Filter == nil {
		searchRequest.Filter = &dto.EventFilter{}
	}
	if searchRequest.Filter.CreatorOrInvitedUserUuid == nil {
		searchRequest.Filter.CreatorOrInvitedUserUuid = &filter.OrQueryFilter{}
	}

	searchRequest.Filter.CreatorOrInvitedUserUuid = &filter.OrQueryFilter{Eq: userUuid}

	return nil
}
