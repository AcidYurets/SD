package integrational

import (
	"calend/internal/models/session"
	"calend/internal/modules/domain/user/dto"
	"context"
	"github.com/google/uuid"
)

func makeCtxByUser(user *dto.User) context.Context {
	ss := session.Session{
		SID:      uuid.NewString(),
		UserUuid: user.Uuid,
	}

	ctx := context.Background()
	return session.SetSessionToCtx(ctx, ss)
}
