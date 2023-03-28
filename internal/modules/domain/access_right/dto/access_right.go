package dto

import "calend/internal/models/access"

type AccessRight struct {
	Code        access.Type // Код
	Description string      // Описание
}

type AccessRights []*AccessRight
