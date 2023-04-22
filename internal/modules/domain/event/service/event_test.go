package service

import (
	"calend/internal/models/session"
	ar_dto "calend/internal/modules/domain/access_right/dto"
	ev_dto "calend/internal/modules/domain/event/dto"
	tag_dto "calend/internal/modules/domain/tag/dto"
	user_dto "calend/internal/modules/domain/user/dto"
	"calend/internal/utils/ptr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	expectedEvent = &ev_dto.Event{
		Uuid:        "123e4567-e89b-12d3-a456-426655440044",
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
					Code:        "riud",
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
	}
)

func ctxWithSession() context.Context {
	ss := session.Session{
		SID:      "509829d8-ab00-4608-a8d9-432bf7f9b118",
		UserUuid: "123e4567-e89b-12d3-a456-426655440000",
	}

	ctx := context.Background()
	return session.SetSessionToCtx(ctx, ss)
}

func TestEventService_GetByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	eRepo := NewMockIEventRepo(ctrl)
	iRepo := NewMockIInvitationRepo(ctrl)
	service := NewEventService(eRepo, iRepo)
	ctx := ctxWithSession()

	uuid := "123e4567-e89b-12d3-a456-426655440044"

	eRepo.EXPECT().GetByUuid(ctx, uuid).Return(expectedEvent, nil).Times(2)

	event, err := service.GetByUuid(ctx, uuid)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvent, event)
}

func TestEventService_ListAvailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	eRepo := NewMockIEventRepo(ctrl)
	iRepo := NewMockIInvitationRepo(ctrl)
	service := NewEventService(eRepo, iRepo)
	ctx := ctxWithSession()

	userUuid := "123e4567-e89b-12d3-a456-426655440000"
	expectedEvents := ev_dto.Events{expectedEvent}

	eRepo.EXPECT().ListAvailable(ctx, userUuid).Return(expectedEvents, nil)

	events, err := service.ListAvailable(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, events)
}

func TestEventService_CreateWithInvitations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	eRepo := NewMockIEventRepo(ctrl)
	iRepo := NewMockIInvitationRepo(ctrl)
	service := NewEventService(eRepo, iRepo)
	ctx := ctxWithSession()

	uuid := "123e4567-e89b-12d3-a456-426655440044"
	createEvent := &ev_dto.CreateEvent{
		Timestamp:   time.Now(),
		Name:        "Событие",
		Description: ptr.String("Описание"),
		Type:        "Мероприятие",
		IsWholeDay:  false,
		TagUuids:    []string{"Тег"},
	}
	createEventInvs := ev_dto.CreateInvitations{
		&ev_dto.CreateInvitation{
			UserUuid:        "123e4567-e89b-12d3-a456-426655440000",
			AccessRightCode: "r",
		},
	}

	createInvs := ev_dto.CreateInvitations{
		&ev_dto.CreateInvitation{
			EventUuid:       uuid,
			UserUuid:        "123e4567-e89b-12d3-a456-426655440000",
			AccessRightCode: "r",
		},
	}

	createdInvs := ev_dto.Invitations{
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
	}
	createdEventWithoutInvs := &ev_dto.Event{
		Uuid:        "123e4567-e89b-12d3-a456-426655440044",
		Timestamp:   time.Now(),
		Name:        "Событие",
		Description: ptr.String("Описание"),
		Type:        "Мероприятие",
		IsWholeDay:  false,
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
	}

	eRepo.EXPECT().Create(ctx, createEvent).Return(createdEventWithoutInvs, nil)
	iRepo.EXPECT().CreateBulk(ctx, createInvs).Return(createdInvs, nil)
	eRepo.EXPECT().GetByUuid(ctx, uuid).Return(expectedEvent, nil)

	event, err := service.Create(ctx, createEvent, createEventInvs)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvent, event)
}

func TestEventService_AddInvitations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	eRepo := NewMockIEventRepo(ctrl)
	iRepo := NewMockIInvitationRepo(ctrl)
	service := NewEventService(eRepo, iRepo)
	ctx := ctxWithSession()

	uuid := "123e4567-e89b-12d3-a456-426655440044"

	createEventInvs := ev_dto.CreateInvitations{
		&ev_dto.CreateInvitation{
			UserUuid:        "123e4567-e89b-12d3-a456-426655440000",
			AccessRightCode: "r",
		},
	}

	createInvs := ev_dto.CreateInvitations{
		&ev_dto.CreateInvitation{
			EventUuid:       uuid,
			UserUuid:        "123e4567-e89b-12d3-a456-426655440000",
			AccessRightCode: "r",
		},
	}

	createdInvs := ev_dto.Invitations{
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
	}

	iRepo.EXPECT().CreateBulk(ctx, createInvs).Return(createdInvs, nil)
	eRepo.EXPECT().GetByUuid(ctx, uuid).Return(expectedEvent, nil).Times(2)

	event, err := service.AddInvitations(ctx, uuid, createEventInvs)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvent, event)
}

func TestEventService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	eRepo := NewMockIEventRepo(ctrl)
	iRepo := NewMockIInvitationRepo(ctrl)
	service := NewEventService(eRepo, iRepo)
	ctx := ctxWithSession()

	uuid := "123e4567-e89b-12d3-a456-426655440044"
	updateEvent := &ev_dto.UpdateEvent{
		Timestamp:   time.Now(),
		Name:        "Событие",
		Description: ptr.String("Описание"),
		Type:        "Мероприятие",
		IsWholeDay:  false,
		TagUuids:    []string{"Тег"},
	}
	createEventInvs := ev_dto.CreateInvitations{
		&ev_dto.CreateInvitation{
			UserUuid:        "123e4567-e89b-12d3-a456-426655440000",
			AccessRightCode: "r",
		},
	}

	createInvs := ev_dto.CreateInvitations{
		&ev_dto.CreateInvitation{
			EventUuid:       uuid,
			UserUuid:        "123e4567-e89b-12d3-a456-426655440000",
			AccessRightCode: "r",
		},
	}
	createdInvs := ev_dto.Invitations{
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
	}
	updatedEvent := &ev_dto.Event{
		Uuid:        "123e4567-e89b-12d3-a456-426655440044",
		Timestamp:   time.Now(),
		Name:        "Событие",
		Description: ptr.String("Описание"),
		Type:        "Мероприятие",
		IsWholeDay:  false,
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
	}

	eRepo.EXPECT().GetByUuid(ctx, uuid).Return(expectedEvent, nil).Times(2)
	eRepo.EXPECT().Update(ctx, uuid, updateEvent).Return(updatedEvent, nil)
	iRepo.EXPECT().DeleteByEventUuid(ctx, uuid).Return(nil)
	iRepo.EXPECT().CreateBulk(ctx, createInvs).Return(createdInvs, nil)

	event, err := service.Update(ctx, uuid, updateEvent, createEventInvs)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvent, event)
}

func TestEventService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	eRepo := NewMockIEventRepo(ctrl)
	iRepo := NewMockIInvitationRepo(ctrl)
	service := NewEventService(eRepo, iRepo)
	ctx := ctxWithSession()

	uuid := "123e4567-e89b-12d3-a456-426655440044"

	eRepo.EXPECT().GetByUuid(ctx, uuid).Return(expectedEvent, nil)
	eRepo.EXPECT().Delete(ctx, uuid).Return(nil)
	iRepo.EXPECT().DeleteByEventUuid(ctx, uuid).Return(nil)

	err := service.Delete(ctx, uuid)
	assert.NoError(t, err)
}
