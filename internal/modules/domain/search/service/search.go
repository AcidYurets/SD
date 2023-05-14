package service

import (
	ev_dto "calend/internal/modules/domain/event/dto"
	"calend/internal/modules/domain/search/dto"
	"context"
)

//go:generate mockgen -destination mock_test.go -package service . ISearchRepo

type ISearchRepo interface {
	SearchEvents(ctx context.Context, searchRequest *dto.EventSearchRequest) (ev_dto.Events, error)
}

type SearchService struct {
	repo ISearchRepo
}

func NewSearchService(repo ISearchRepo) *SearchService {
	return &SearchService{
		repo: repo,
	}
}

func (r *SearchService) SearchEvents(ctx context.Context, searchRequest *dto.EventSearchRequest) (ev_dto.Events, error) {
	// defer timer.Evaluate("SearchEvents")()
	return r.repo.SearchEvents(ctx, searchRequest)
}
