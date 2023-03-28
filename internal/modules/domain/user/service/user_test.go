package service

import (
	"calend/internal/modules/domain/user/dto"
	"calend/internal/utils/ptr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIUserRepo(ctrl)
	service := NewUserService(repo)

	uuid := "123e4567-e89b-12d3-a456-426655440000"

	repo.EXPECT().Delete(gomock.Any(), uuid).Return(nil)

	err := service.Delete(context.Background(), uuid)
	assert.NoError(t, err)
}

func TestUserService_GetByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIUserRepo(ctrl)
	service := NewUserService(repo)

	uuid := "123e4567-e89b-12d3-a456-426655440000"
	expectedUser := &dto.User{
		Uuid:         uuid,
		Phone:        "89197628803",
		Login:        "user",
		PasswordHash: "1234",
	}

	repo.EXPECT().GetByUuid(gomock.Any(), uuid).Return(expectedUser, nil)

	user, err := service.GetByUuid(context.Background(), uuid)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIUserRepo(ctrl)
	service := NewUserService(repo)

	expectedUsers := dto.Users{
		&dto.User{
			Uuid:         "123e4567-e89b-12d3-a456-426655440000",
			Phone:        "89197628803",
			Login:        "user1",
			PasswordHash: "1234",
		},
		&dto.User{
			Uuid:         "123e4567-e89b-12d3-a456-426655440001",
			Phone:        "89197628803",
			Login:        "user2",
			PasswordHash: "12345",
		},
	}

	repo.EXPECT().List(gomock.Any()).Return(expectedUsers, nil)

	users, err := service.List(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestUserService_Restore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIUserRepo(ctrl)
	service := NewUserService(repo)

	uuid := "123e4567-e89b-12d3-a456-426655440000"
	expectedUser := &dto.User{
		Uuid:         uuid,
		Phone:        "89197628803",
		Login:        "user",
		PasswordHash: "1234",
	}

	repo.EXPECT().Restore(gomock.Any(), uuid).Return(expectedUser, nil)

	user, err := service.Restore(context.Background(), uuid)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIUserRepo(ctrl)
	service := NewUserService(repo)

	uuid := "123e4567-e89b-12d3-a456-426655440000"
	updateUser := &dto.UpdateUser{
		Phone: ptr.String("89197628804"),
		Login: ptr.String("user2"),
	}

	expectedUser := &dto.User{
		Uuid:         uuid,
		Phone:        "89197628804",
		Login:        "user2",
		PasswordHash: "1234",
	}

	repo.EXPECT().Update(gomock.Any(), uuid, updateUser).Return(expectedUser, nil)

	user, err := service.Update(context.Background(), uuid, updateUser)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}
