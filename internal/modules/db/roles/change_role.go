package roles

import (
	"calend/internal/models/roles"
	"calend/internal/models/session"
	"context"
	entsql "entgo.io/ent/dialect/sql"
)

func ChangeRole(ctx context.Context, conn entsql.ExecQuerier) error {
	var prepQuery string
	args := make([]any, 0)

	if roles.CheckNeedChange(ctx) {
		currentUserUuid, err := session.GetUserUuidFromCtx(ctx)
		if err != nil {
			return err
		}

		prepQuery = "CALL before_each_query($1);\n"
		args = append(args, currentUserUuid)
	}
	if roles.CheckUseSuperUser(ctx) {
		prepQuery = "SET ROLE superuser;\n"
	}

	_, err := conn.ExecContext(ctx, prepQuery, args...)
	if err != nil {
		return err
	}

	return nil
}
