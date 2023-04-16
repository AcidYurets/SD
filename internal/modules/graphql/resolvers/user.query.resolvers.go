package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"calend/internal/modules/domain/user/dto"
	"context"
)

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, id string) (*dto.User, error) {
	user, err := r.userService.GetByUuid(ctx, id)

	return user, err
}
