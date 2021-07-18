package graphql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"meetmeup/graphql/generated"
	"meetmeup/model"
)

type Resolver struct{}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	panic("not implemented")
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	panic("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
