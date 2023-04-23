package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"coffeewithegg/apps/adam/graph/generated"
	"coffeewithegg/apps/adam/graph/model"
	"coffeewithegg/apps/adam/project/service"
	"coffeewithegg/apps/adam/utils"
	"context"
)

// UpsertProject is the resolver for the upsertProject field.
func (r *mutationResolver) UpsertProject(ctx context.Context, input *model.ProjectInput) (*model.Project, error) {
	var projectService *service.ProjectService
	err := utils.ResolveService(projectService.GetServiceName(), &projectService)
	if err != nil {
		return nil, err
	}
	return projectService.UpsertProject(ctx, input)
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, filters *model.ProjectFilter) ([]*model.Project, error) {
	var projectService *service.ProjectService
	err := utils.ResolveService(projectService.GetServiceName(), &projectService)
	if err != nil {
		return nil, err
	}
	return projectService.GetProjects(ctx, filters)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
