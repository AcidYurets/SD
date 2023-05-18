package dataloader

import (
	ar_serv "calend/internal/modules/domain/access_right/service"
	auth_serv "calend/internal/modules/domain/auth/service"
	event_serv "calend/internal/modules/domain/event/service"
	tag_serv "calend/internal/modules/domain/tag/service"
	user_serv "calend/internal/modules/domain/user/service"
	"github.com/graph-gophers/dataloader"
)

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	AccessRightLoader *dataloader.Loader
	UserLoader        *dataloader.Loader
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(
	userService *user_serv.UserService,
	tagService *tag_serv.TagService,
	authService *auth_serv.AuthService,
	arService *ar_serv.AccessRightService,
	eventService *event_serv.EventService,
) *Loaders {
	serverLoader := &ServerLoader{
		userService:  userService,
		tagService:   tagService,
		authService:  authService,
		arService:    arService,
		eventService: eventService,
	}

	loaders := &Loaders{
		AccessRightLoader: dataloader.NewBatchedLoader(serverLoader.GetAccessRightsByCodes,
			dataloader.WithClearCacheOnBatch()),
		UserLoader: dataloader.NewBatchedLoader(serverLoader.GetUsersByUuids,
			dataloader.WithClearCacheOnBatch()),
	}
	return loaders
}
