package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"coffeewithegg/apps/adam/graph/generated"
	"coffeewithegg/apps/adam/graph/model"
	"coffeewithegg/apps/adam/project/service"
	"context"
)

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, filters *model.ProjectFilter) ([]*model.Project, error) {
	return projectService.GetProjects(ctx, filters)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

var projectService = service.NewProjectService()
