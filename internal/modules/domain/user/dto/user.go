package dto

type User struct {
	Uuid         string // Uuid пользователя
	Phone        string // Телефон
	Login        string // Логин в системе
	PasswordHash string // Хэш пароля
}

type Users []*User

type UpdateUser struct {
	Phone        *string // Телефон
	Login        *string // Логин в системе
	PasswordHash *string // Хэш пароля
}
