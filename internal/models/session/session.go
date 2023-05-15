package session

import (
	"calend/internal/models/roles"
	"context"
	"fmt"
)

// Session сессия пользователя
type Session struct {
	SID      string     // Уникальный идентификатор сессии
	UserUuid string     // Uuid пользователя
	Role     roles.Type // Роль пользователя
}

type sessionCtx struct{}

func GetSessionFromCtx(ctx context.Context) (session Session, ok bool) {
	session, ok = ctx.Value(sessionCtx{}).(Session)
	return session, ok
}

func SetSessionToCtx(ctx context.Context, session Session) context.Context {
	return context.WithValue(ctx, sessionCtx{}, session)
}

func GetUserUuidFromCtx(ctx context.Context) (string, error) {
	s, ok := GetSessionFromCtx(ctx)
	if !ok {
		return "", fmt.Errorf("cессия отсутствует в контексте")
	}

	return s.UserUuid, nil
}
