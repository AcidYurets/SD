package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"calend/internal/utils/ptr"
	"context"
)

// GenerateUsers is the resolver for the GenerateUsers field.
func (r *mutationResolver) GenerateUsers(ctx context.Context, count uint) (*string, error) {
	err := r.userGenerator.Generate(ctx, count)
	if err != nil {
		return nil, err
	}

	return ptr.String("Готово"), nil
}

// GenerateTags is the resolver for the GenerateTags field.
func (r *mutationResolver) GenerateTags(ctx context.Context, count uint) (*string, error) {
	err := r.tagGenerator.Generate(ctx, count)
	if err != nil {
		return nil, err
	}

	return ptr.String("Готово"), nil
}

// GenerateEvents is the resolver for the GenerateEvents field.
func (r *mutationResolver) GenerateEvents(ctx context.Context, count uint) (*string, error) {
	err := r.eventGenerator.Generate(ctx, count)
	if err != nil {
		return nil, err
	}

	return ptr.String("Готово"), nil
}
