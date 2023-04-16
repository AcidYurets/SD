package dto

import (
	"calend/internal/models/session"
	"github.com/golang-jwt/jwt/v4"
)

// UserCredentials учетные данные пользователя
type UserCredentials struct {
	Login    string // Логин в системе
	Password string // Пароль
}

// NewUser данные нового пользователя
type NewUser struct {
	Login    string // Логин в системе
	Password string // Пароль
	Phone    string // Телефон
}

// Tokens токены
type Tokens struct {
	AccessToken  string // Токен доступа
	RefreshToken string // Токен обновления токена доступа
}

// JWT ответ на успешную аутентификацию
type JWT struct {
	Tokens
	Session *session.Session // Сессия пользователя
}

// TokenClaims тип данных токена
type TokenClaims struct {
	jwt.RegisteredClaims
	Session *session.Session `json:"session"` // Сессия в данных токена
}
