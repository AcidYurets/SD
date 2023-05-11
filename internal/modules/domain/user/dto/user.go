package dto

import "calend/internal/models/roles"

type User struct {
	Uuid         string     // Uuid пользователя
	Phone        string     // Телефон
	Login        string     // Логин в системе
	PasswordHash string     // Хэш пароля
	Role         roles.Type // Роль пользователя
}

type Users []*User

type CreateUser struct {
	Phone        string     // Телефон
	Login        string     // Логин в системе
	PasswordHash string     // Хэш пароля
	Role         roles.Type // Роль пользователя
}

type UpdateUser struct {
	Phone *string     // Телефон
	Login *string     // Логин в системе
	Role  *roles.Type // Роль пользователя
}
