package integrational

import (
	"calend/internal/modules/db/ent"
	ar_serv "calend/internal/modules/domain/access_right/service"
	auth_serv "calend/internal/modules/domain/auth/service"
	event_serv "calend/internal/modules/domain/event/service"
	tag_serv "calend/internal/modules/domain/tag/service"
	user_serv "calend/internal/modules/domain/user/service"
	"context"
	"go.uber.org/fx"
	"testing"
)

func TestServices(t *testing.T) {
	fx.New(
		testModule,
		testInvokables,

		fx.Supply(t),
		fx.Invoke(execTests),
	).Run()

}

func execTests(
	t *testing.T,
	userService *user_serv.UserService,
	tagService *tag_serv.TagService,
	authService *auth_serv.AuthService,
	arService *ar_serv.AccessRightService,
	eventService *event_serv.EventService,

	client *ent.Client,
	lifecycle fx.Lifecycle,
	shutdowner fx.Shutdowner,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				eventServiceTest(t, eventService, tagService, authService, client)
				authServiceTest(t, userService, authService, client)
				userServiceTest(t, userService, authService, client)
				tagServiceTest(t, tagService, authService, client)
				accessRightServiceTest(t, arService, authService, client)

				_ = shutdowner.Shutdown()
			}()

			return nil
		},
	})
}
