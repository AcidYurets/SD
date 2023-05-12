package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	dto3 "calend/internal/modules/domain/access_right/dto"
	"calend/internal/modules/domain/event/dto"
	dto1 "calend/internal/modules/domain/tag/dto"
	dto2 "calend/internal/modules/domain/user/dto"
	"calend/internal/modules/graphql/generated"
	"context"
)

// Tags is the resolver for the Tags field.
func (r *eventResolver) Tags(ctx context.Context, obj *dto.Event) ([]*dto1.Tag, error) {
	tags, err := r.eventService.ListTagsByEventUuid(ctx, obj.Uuid)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// Creator is the resolver for the Creator field.
func (r *eventResolver) Creator(ctx context.Context, obj *dto.Event) (*dto2.User, error) {
	creator, err := r.userService.GetByUuid(ctx, obj.CreatorUuid)
	if err != nil {
		return nil, err
	}

	return creator, nil
}

// User is the resolver for the User field.
func (r *invitationResolver) User(ctx context.Context, obj *dto.Invitation) (*dto2.User, error) {
	user, err := r.userService.GetByUuid(ctx, obj.UserUuid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// AccessRight is the resolver for the AccessRight field.
func (r *invitationResolver) AccessRight(ctx context.Context, obj *dto.Invitation) (*dto3.AccessRight, error) {
	accessRight, err := r.arService.GetByCode(ctx, obj.AccessRightCode)
	if err != nil {
		return nil, err
	}

	return accessRight, nil
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

// Invitation returns generated.InvitationResolver implementation.
func (r *Resolver) Invitation() generated.InvitationResolver { return &invitationResolver{r} }

type eventResolver struct{ *Resolver }
type invitationResolver struct{ *Resolver }