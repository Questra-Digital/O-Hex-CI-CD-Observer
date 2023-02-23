package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"gitlab/jobs/graph/model"
)

// CreateJob is the resolver for the CreateJob field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.CreateJobInput) (*model.Jobs, error) {
	panic(fmt.Errorf("not implemented: CreateJob - CreateJob"))
}

// JobsList is the resolver for the jobsList field.
func (r *queryResolver) JobsList(ctx context.Context) ([]*model.Jobs, error) {
	panic(fmt.Errorf("not implemented: JobsList - jobsList"))
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*model.Jobs, error) {
	panic(fmt.Errorf("not implemented: Job - job"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
