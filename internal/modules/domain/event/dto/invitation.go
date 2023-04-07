package dto

import (
	"calend/internal/models/access"
	ar_dto "calend/internal/modules/domain/access_right/dto"
	user_dto "calend/internal/modules/domain/user/dto"
)

type Invitation struct {
	Uuid        string              // Uuid приглашения
	User        *user_dto.User      // Приглашенный пользователь
	AccessRight *ar_dto.AccessRight // Права доступа приглашенного пользователя
}

type Invitations []*Invitation

// CreateInvitation модель для создания приглашений
type CreateInvitation struct {
	EventUuid       string      // Uuid события
	UserUuid        string      // Uuid приглашенного пользователя
	AccessRightCode access.Type // Код права доступа
}

type CreateInvitations []*CreateInvitation
