package session

import (
	"context"
)

// Session сессия пользователя
type Session struct {
	SID      string // Уникальный идентификатор сессии
	UserUuid string // Uuid пользователя
}

type sessionCtx struct{}

func GetSessionFromCtx(ctx context.Context) (session Session, ok bool) {
	session, ok = ctx.Value(sessionCtx{}).(Session)
	return session, ok
}

func SetSessionToCtx(ctx context.Context, session Session) context.Context {
	return context.WithValue(ctx, sessionCtx{}, session)
}
