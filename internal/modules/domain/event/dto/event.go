package dto

import (
	"time"
)

type Event struct {
	Uuid        string        // Uuid события
	Timestamp   time.Time     // Временная метка
	Name        string        // Название
	Description *string       // Описание
	Type        string        // Тип события
	IsWholeDay  bool          // Событие на целый день?
	Invitations []*Invitation // Приглашения события
	CreatorUuid string        // Создатель события
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
