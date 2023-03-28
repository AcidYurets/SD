package dto

type AccessRight struct {
	Code        AccessCode // Код
	Description string     // Описание
}

type AccessRights []*AccessRight
