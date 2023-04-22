package seed

import (
	"calend/internal/models/access"
	"calend/internal/modules/db/ent"
	"context"
)

func AccessRight(client *ent.Client) error {
	var err error
	ctx := context.Background()

	_, err = client.AccessRight.Create().SetID(access.OnlyReadAccess).SetDescription("Право на просмотр").Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.AccessRight.Create().SetID(access.ReadInviteAccess).SetDescription("Право на просмотр и приглашение").Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.AccessRight.Create().SetID(access.ReadInviteUpdateAccess).SetDescription("Право на просмотр, приглашение и изменение").Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.AccessRight.Create().SetID(access.FullAccess).SetDescription("Право на просмотр, приглашение, изменение и удаление").Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
