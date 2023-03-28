package dto

import (
	inv_dto "calend/internal/modules/domain/invitation/dto"
	tag_dto "calend/internal/modules/domain/tag/dto"
	user_dto "calend/internal/modules/domain/user/dto"
	"time"
)

type Event struct {
	Uuid        string              // Uuid пользователя
	Timestamp   time.Time           // Временная метка
	Name        string              // Название
	Description *string             // Описание
	Type        string              // Тип события
	IsWholeDay  bool                // Событие на целый день?
	Invitations inv_dto.Invitations // Приглашения события
	Tags        tag_dto.Tags        // Теги события
	Creator     user_dto.User       // Создатель события
}

type Events []*Event

// CreateEvent модель для создания события без приглашений
type CreateEvent struct {
	Uuid        string    // Uuid пользователя
	Timestamp   time.Time // Временная метка
	Name        string    // Название
	Description *string   // Описание
	Type        string    // Тип события
	IsWholeDay  bool      // Событие на целый день?
	TagUuids    []string  // Массив Uuid-ов тегов события
	CreatorUuid string    // Создатель события
}

// UpdateEvent модель для создания события без приглашений
type UpdateEvent struct {
	Timestamp   *time.Time // Временная метка
	Name        *string    // Название
	Description *string    // Описание
	Type        *string    // Тип события
	IsWholeDay  *bool      // Событие на целый день?
	TagUuids    []string   // Измененный массив Uuid-ов тегов события
}

// CreateEventInvitation модель для создания приглашений вмести с событием
type CreateEventInvitation struct {
	UserUuid        string // Uuid приглашенного пользователя
	AccessRightCode string // Код права доступа
}

type CreateEventInvitations []*CreateEventInvitation
