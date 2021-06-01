package Gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"MyGqlgen/generated"
	"MyGqlgen/model"
	"context"
)

type Resolver struct{}

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetups, error) {
	panic("not implemented")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
