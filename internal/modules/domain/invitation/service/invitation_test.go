package service

import (
	ar_dto "calend/internal/modules/domain/access_right/dto"
	"calend/internal/modules/domain/invitation/dto"
	user_dto "calend/internal/modules/domain/user/dto"
	"calend/internal/utils/ptr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvitationService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIInvitationRepo(ctrl)
	service := NewInvitationService(repo)

	uuid := "123e4567-e89b-12d3-a456-426655440000"

	repo.EXPECT().Delete(gomock.Any(), uuid).Return(nil)

	err := service.Delete(context.Background(), uuid)
	assert.NoError(t, err)
}

func TestInvitationService_GetByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIInvitationRepo(ctrl)
	service := NewInvitationService(repo)

	uuid := "123e4567-e89b-12d3-a456-426655440000"
	expectedInvitation := &dto.Invitation{
		Uuid: uuid,
		User: &user_dto.User{
			Uuid:         "123e4567-e89b-12d3-a456-426655440000",
			Phone:        "89197628803",
			Login:        "user1",
			PasswordHash: "1234",
		},
		AccessRight: &ar_dto.AccessRight{
			Code:        "r",
			Description: "Право только на чтение информации о событии",
		},
	}

	repo.EXPECT().GetByUuid(gomock.Any(), uuid).Return(expectedInvitation, nil)

	user, err := service.GetByUuid(context.Background(), uuid)
	assert.NoError(t, err)
	assert.Equal(t, expectedInvitation, user)
}

func TestInvitationService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIInvitationRepo(ctrl)
	service := NewInvitationService(repo)

	expectedInvitations := dto.Invitations{
		&dto.Invitation{
			Uuid: "123e4567-e89b-12d3-a456-426655440000",
			User: &user_dto.User{
				Uuid:         "123e4567-e89b-12d3-a456-426655440000",
				Phone:        "89197628803",
				Login:        "user1",
				PasswordHash: "1234",
			},
			AccessRight: &ar_dto.AccessRight{
				Code:        "r",
				Description: "Право только на чтение информации о событии",
			},
		},
	}

	repo.EXPECT().List(gomock.Any()).Return(expectedInvitations, nil)

	users, err := service.List(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedInvitations, users)
}

func TestInvitationService_ListByUserUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIInvitationRepo(ctrl)
	service := NewInvitationService(repo)

	userUuid := "123e4567-e89b-12d3-a456-426655440000"
	expectedInvitations := dto.Invitations{
		&dto.Invitation{
			Uuid: "123e4567-e89b-12d3-a456-426655440000",
			User: &user_dto.User{
				Uuid:         userUuid,
				Phone:        "89197628803",
				Login:        "user1",
				PasswordHash: "1234",
			},
			AccessRight: &ar_dto.AccessRight{
				Code:        "r",
				Description: "Право только на чтение информации о событии",
			},
		},
	}

	repo.EXPECT().ListByUserUuid(gomock.Any(), userUuid).Return(expectedInvitations, nil)

	users, err := service.ListByUserUuid(context.Background(), userUuid)
	assert.NoError(t, err)
	assert.Equal(t, expectedInvitations, users)
}

func TestInvitationService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIInvitationRepo(ctrl)
	service := NewInvitationService(repo)

	user1Uuid := "123e4567-e89b-12d3-a456-426655440000"
	user2Uuid := "123e4567-e89b-12d3-a456-426655440001"
	eventUuid := "123e4567-e89b-12d3-a456-426655440000"
	ar := "r"
	createInvitations := dto.CreateInvitations{
		&dto.CreateInvitation{
			UserUuid:        user1Uuid,
			EventUuid:       eventUuid,
			AccessRightCode: ar,
		},
		&dto.CreateInvitation{
			UserUuid:        user2Uuid,
			EventUuid:       eventUuid,
			AccessRightCode: ar,
		},
	}

	expectedInvitations := dto.Invitations{
		&dto.Invitation{
			Uuid: "123e4567-e89b-12d3-a456-426655440000",
			User: &user_dto.User{
				Uuid:         user1Uuid,
				Phone:        "89197628803",
				Login:        "user1",
				PasswordHash: "1234",
			},
			AccessRight: &ar_dto.AccessRight{
				Code:        ar,
				Description: "Право только на чтение информации о событии",
			},
		},
		&dto.Invitation{
			Uuid: "123e4567-e89b-12d3-a456-426655440001",
			User: &user_dto.User{
				Uuid:         user2Uuid,
				Phone:        "89197628803",
				Login:        "user1",
				PasswordHash: "1234",
			},
			AccessRight: &ar_dto.AccessRight{
				Code:        ar,
				Description: "Право только на чтение информации о событии",
			},
		},
	}

	repo.EXPECT().CreateBulk(gomock.Any(), createInvitations).Return(expectedInvitations, nil)

	invs, err := service.CreateBulk(context.Background(), createInvitations)
	assert.NoError(t, err)
	assert.Equal(t, expectedInvitations, invs)
}

func TestInvitationService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIInvitationRepo(ctrl)
	service := NewInvitationService(repo)

	uuid := "123e4567-e89b-12d3-a456-426655440000"
	updateInvitation := &dto.UpdateInvitation{
		AccessRightCode: ptr.String("ri"),
	}

	expectedInvitation := &dto.Invitation{
		Uuid: uuid,
		User: &user_dto.User{
			Uuid:         "123e4567-e89b-12d3-a456-426655440000",
			Phone:        "89197628803",
			Login:        "user1",
			PasswordHash: "1234",
		},
		AccessRight: &ar_dto.AccessRight{
			Code:        "ri",
			Description: "Право только на чтение информации о событии",
		},
	}

	repo.EXPECT().Update(gomock.Any(), uuid, updateInvitation).Return(expectedInvitation, nil)

	user, err := service.Update(context.Background(), uuid, updateInvitation)
	assert.NoError(t, err)
	assert.Equal(t, expectedInvitation, user)
}
