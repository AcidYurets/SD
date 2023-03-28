package service

import (
	"calend/internal/models/access"
	"calend/internal/modules/domain/access_right/dto"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessRightService_GetByCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIAccessRightRepo(ctrl)
	service := NewAccessRightService(repo)

	code := access.Type("r")
	expectedAccessRight := &dto.AccessRight{
		Code:        code,
		Description: "Право только на чтение информации о событии",
	}

	repo.EXPECT().GetByCode(gomock.Any(), code).Return(expectedAccessRight, nil)

	user, err := service.GetByCode(context.Background(), code)
	assert.NoError(t, err)
	assert.Equal(t, expectedAccessRight, user)
}

func TestAccessRightService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIAccessRightRepo(ctrl)
	service := NewAccessRightService(repo)

	expectedAccessRights := dto.AccessRights{
		&dto.AccessRight{
			Code:        access.Type("r"),
			Description: "Право только на чтение информации о событии",
		},
		&dto.AccessRight{
			Code:        access.Type("ri"),
			Description: "Право на чтение информации о событии и на приглашение других пользователей",
		},
	}

	repo.EXPECT().List(gomock.Any()).Return(expectedAccessRights, nil)

	users, err := service.List(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedAccessRights, users)
}
