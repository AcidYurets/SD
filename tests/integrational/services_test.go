package integrational

import (
	"context"
	"go.uber.org/fx"
	"testing"
)

func TestServices(t *testing.T) {
	fx.New(
		TestModule,
		TestInvokables,

		fx.Supply(t),
		fx.Invoke(UserServiceTest),

		fx.Invoke(shutdownAfterTests),
	).Run()

}

func shutdownAfterTests(lifecycle fx.Lifecycle, shutdowner fx.Shutdowner) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := shutdowner.Shutdown()
			return err
		},
	})
}
