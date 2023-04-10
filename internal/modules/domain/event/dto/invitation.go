package dto

import (
	"calend/internal/models/access"
)

type Invitation struct {
	Uuid            string      // Uuid приглашения
	UserUuid        string      // Приглашенный пользователь
	AccessRightCode access.Type // Права доступа приглашенного пользователя
}

type Invitations []*Invitation

// CreateInvitation модель для создания приглашений
type CreateInvitation struct {
	EventUuid       string      // Uuid события
	UserUuid        string      // Uuid приглашенного пользователя
	AccessRightCode access.Type // Код права доступа
}

type CreateInvitations []*CreateInvitation
