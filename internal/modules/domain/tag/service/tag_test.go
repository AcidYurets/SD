package service

import (
	dto "calend/internal/modules/domain/tag/dto"
	"calend/internal/utils/ptr"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockITagRepo(ctrl)
	service := NewTagService(repo)

	uuid := "123e4567-e89b-12d3-a456-426655440000"

	repo.EXPECT().Delete(gomock.Any(), uuid).Return(nil)

	err := service.Delete(context.Background(), uuid)
	assert.NoError(t, err)
}

func TestTagService_GetByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockITagRepo(ctrl)
	service := NewTagService(repo)

	uuid := "123e4567-e89b-12d3-a456-426655440000"
	expectedTag := &dto.Tag{
		Uuid:        uuid,
		Name:        "Праздник",
		Description: "Тег для праздников",
	}

	repo.EXPECT().GetByUuid(gomock.Any(), uuid).Return(expectedTag, nil)

	user, err := service.GetByUuid(context.Background(), uuid)
	assert.NoError(t, err)
	assert.Equal(t, expectedTag, user)
}

func TestTagService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockITagRepo(ctrl)
	service := NewTagService(repo)

	expectedTags := dto.Tags{
		&dto.Tag{
			Uuid:        "123e4567-e89b-12d3-a456-426655440000",
			Name:        "Праздник",
			Description: "Тег для праздников",
		},
		&dto.Tag{
			Uuid:        "123e4567-e89b-12d3-a456-426655440001",
			Name:        "ДР",
			Description: "Тег для дня рождения",
		},
	}

	repo.EXPECT().List(gomock.Any()).Return(expectedTags, nil)

	users, err := service.List(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedTags, users)
}

func TestTagService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockITagRepo(ctrl)
	service := NewTagService(repo)

	createTag := &dto.CreateTag{
		Name:        "Праздник",
		Description: "Тег для праздников",
	}

	expectedTag := &dto.Tag{
		Uuid:        "123e4567-e89b-12d3-a456-426655440000",
		Name:        "Праздник",
		Description: "Тег для праздников",
	}

	repo.EXPECT().Create(gomock.Any(), createTag).Return(expectedTag, nil)

	user, err := service.Create(context.Background(), createTag)
	assert.NoError(t, err)
	assert.Equal(t, expectedTag, user)
}

func TestTagService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockITagRepo(ctrl)
	service := NewTagService(repo)

	uuid := "123e4567-e89b-12d3-a456-426655440000"
	updateTag := &dto.UpdateTag{
		Description: ptr.String("Тег для праздников (новое описание)"),
	}

	expectedTag := &dto.Tag{
		Uuid:        uuid,
		Name:        "Праздник",
		Description: "Тег для праздников (новое описание)",
	}

	repo.EXPECT().Update(gomock.Any(), uuid, updateTag).Return(expectedTag, nil)

	user, err := service.Update(context.Background(), uuid, updateTag)
	assert.NoError(t, err)
	assert.Equal(t, expectedTag, user)
}
