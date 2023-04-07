package dto

import (
	"calend/internal/models/access"
	tag_dto "calend/internal/modules/domain/tag/dto"
	user_dto "calend/internal/modules/domain/user/dto"
	"time"
)

type Event struct {
	Uuid        string         // Uuid события
	Timestamp   time.Time      // Временная метка
	Name        string         // Название
	Description *string        // Описание
	Type        string         // Тип события
	IsWholeDay  bool           // Событие на целый день?
	Invitations Invitations    // Приглашения события
	Tags        tag_dto.Tags   // Теги события
	Creator     *user_dto.User // Создатель события
}

type Events []*Event

// CreateEvent модель для создания события без приглашений
type CreateEvent struct {
	Timestamp   time.Time // Временная метка
	Name        string    // Название
	Description *string   // Описание
	Type        string    // Тип события
	IsWholeDay  bool      // Событие на целый день?
	TagUuids    []string  // Массив Uuid-ов тегов события
	CreatorUuid string    // Uuid создателя события
}

// UpdateEvent модель для создания события без приглашений
type UpdateEvent struct {
	Timestamp   time.Time // Временная метка
	Name        string    // Название
	Description *string   // Описание
	Type        string    // Тип события
	IsWholeDay  bool      // Событие на целый день?
	TagUuids    []string  // Измененный массив Uuid-ов тегов события
}

// CreateEventInvitation модель для создания приглашений вмести с событием
type CreateEventInvitation struct {
	UserUuid        string      // Uuid приглашенного пользователя
	AccessRightCode access.Type // Код права доступа
}

type CreateEventInvitations []*CreateEventInvitation
