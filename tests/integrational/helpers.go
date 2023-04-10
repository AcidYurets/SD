package integrational

import (
	"calend/internal/models/session"
	"calend/internal/modules/db/ent"
	"calend/internal/modules/db/schema"
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

func truncateAll(client *ent.Client) error {
	_, err := client.Invitation.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	if err != nil {
		return err
	}

	_, err = client.Event.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	if err != nil {
		return err
	}

	_, err = client.Tag.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	if err != nil {
		return err
	}

	_, err = client.AccessRight.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	if err != nil {
		return err
	}

	_, err = client.User.Delete().Exec(schema.SkipSoftDelete(context.Background()))
	if err != nil {
		return err
	}

	return nil
}

func createAccessRight(client *ent.Client) error {
	var err error

	_, err = client.AccessRight.Create().SetID("r").SetDescription("Право на просмотр").Save(context.Background())
	if err != nil {
		return err
	}
	_, err = client.AccessRight.Create().SetID("ri").SetDescription("Право на просмотр и приглашение").Save(context.Background())
	if err != nil {
		return err
	}
	_, err = client.AccessRight.Create().SetID("riu").SetDescription("Право на просмотр, приглашение и изменение").Save(context.Background())
	if err != nil {
		return err
	}
	_, err = client.AccessRight.Create().SetID("riud").SetDescription("Право на просмотр, приглашение, изменение и удаление").Save(context.Background())
	if err != nil {
		return err
	}

	return nil
}
