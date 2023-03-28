package service

import (
	ar_dto "calend/internal/modules/domain/access_right/dto"
	ev_dto "calend/internal/modules/domain/event/dto"
	"calend/internal/modules/domain/search/dto"
	tag_dto "calend/internal/modules/domain/tag/dto"
	user_dto "calend/internal/modules/domain/user/dto"
	"calend/internal/pkg/search/filter"
	"calend/internal/pkg/search/sort"
	"calend/internal/utils/ptr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSearchService_SearchEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockISearchRepo(ctrl)
	service := NewSearchService(repo)

	sortDirection := sort.DirectionAsc
	searchRequest := &dto.EventSearchRequest{
		Filter: &dto.EventFilter{
			Name: &filter.TextQueryFilter{
				Ts: ptr.String("Сделать ла"),
			},
		},
		Sort: &dto.EventSort{
			Timestamp: &sortDirection,
		},
	}

	expectedEvents := ev_dto.Events{
		&ev_dto.Event{
			Uuid:        "123e4567-e89b-12d3-a456-426655440000",
			Timestamp:   time.Now(),
			Name:        "Событие",
			Description: ptr.String("Описание"),
			Type:        "Мероприятие",
			IsWholeDay:  false,
			Invitations: ev_dto.Invitations{
				&ev_dto.Invitation{
					Uuid: "123e4567-e89b-12d3-a456-426655440000",
					User: &user_dto.User{
						Uuid:         "123e4567-e89b-12d3-a456-426655440000",
						Phone:        "89197628803",
						Login:        "user",
						PasswordHash: "1234",
					},
					AccessRight: &ar_dto.AccessRight{
						Code:        "r",
						Description: "Право на чтение",
					},
				},
			},
			Tags: tag_dto.Tags{
				&tag_dto.Tag{
					Uuid:        "Тег",
					Name:        "Название",
					Description: "Описание",
				},
			},
			Creator: &user_dto.User{
				Uuid:         "123e4567-e89b-12d3-a456-426655440001",
				Phone:        "89197628804",
				Login:        "user1",
				PasswordHash: "1235",
			},
		},
	}

	repo.EXPECT().SearchEvents(gomock.Any(), searchRequest).Return(expectedEvents, nil)

	events, err := service.SearchEvents(context.Background(), searchRequest)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, events)
}
