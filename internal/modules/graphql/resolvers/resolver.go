package resolvers

import (
	ar_serv "calend/internal/modules/domain/access_right/service"
	auth_serv "calend/internal/modules/domain/auth/service"
	event_serv "calend/internal/modules/domain/event/service"
	search_serv "calend/internal/modules/domain/search/service"
	tag_serv "calend/internal/modules/domain/tag/service"
	user_serv "calend/internal/modules/domain/user/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userService   *user_serv.UserService
	tagService    *tag_serv.TagService
	authService   *auth_serv.AuthService
	arService     *ar_serv.AccessRightService
	eventService  *event_serv.EventService
	searchService *search_serv.SearchService
}

func NewResolver(
	userService *user_serv.UserService,
	tagService *tag_serv.TagService,
	authService *auth_serv.AuthService,
	arService *ar_serv.AccessRightService,
	eventService *event_serv.EventService,
	searchService *search_serv.SearchService,
) *Resolver {

	r := &Resolver{
		userService:   userService,
		tagService:    tagService,
		authService:   authService,
		arService:     arService,
		eventService:  eventService,
		searchService: searchService,
	}

	return r
}
