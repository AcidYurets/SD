package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"calend/internal/modules/domain/event/dto"
	"calend/internal/modules/elastic/reindex"
	"context"
)

// EventReindex is the resolver for the EventReindex field.
func (r *mutationResolver) EventReindex(ctx context.Context) (*reindex.Stats, error) {
	stats, err := r.reindexEventService.ReindexAll(ctx)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// EventCreate is the resolver for the EventCreate field.
func (r *mutationResolver) EventCreate(ctx context.Context, event dto.CreateEvent, invitations []*dto.CreateInvitation) (*dto.Event, error) {
	ev, err := r.eventService.Create(ctx, &event, invitations)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

// EventAddInvitations is the resolver for the EventAddInvitations field.
func (r *mutationResolver) EventAddInvitations(ctx context.Context, id string, invitations []*dto.CreateInvitation) (*dto.Event, error) {
	ev, err := r.eventService.AddInvitations(ctx, id, invitations)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

// EventUpdate is the resolver for the EventUpdate field.
func (r *mutationResolver) EventUpdate(ctx context.Context, id string, event dto.UpdateEvent, invitations []*dto.CreateInvitation) (*dto.Event, error) {
	ev, err := r.eventService.Update(ctx, id, &event, invitations)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

// EventDelete is the resolver for the EventDelete field.
func (r *mutationResolver) EventDelete(ctx context.Context, id string) (string, error) {
	err := r.eventService.Delete(ctx, id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// Event is the resolver for the Event field.
func (r *queryResolver) Event(ctx context.Context, id string) (*dto.Event, error) {
	event, err := r.eventService.GetByUuid(ctx, id)
	if err != nil {
		return nil, err
	}

	return event, nil
}

// EventsAvailable is the resolver for the EventsAvailable field.
func (r *queryResolver) EventsAvailable(ctx context.Context) ([]*dto.Event, error) {
	events, err := r.eventService.ListAvailable(ctx)
	if err != nil {
		return nil, err
	}

	return events, nil
}
