package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"coffeewithegg/apps/adam/graph/model"
	"coffeewithegg/apps/adam/infrastructure/service"
	"coffeewithegg/apps/adam/utils"
	"context"
	"fmt"
)

// RequestService is the resolver for the requestService field.
func (r *mutationResolver) RequestService(ctx context.Context, input *model.RequestServiceInput) (*model.Infrastructure, error) {
	var infraService *service.InfrastructureService
	err := utils.ResolveService(infraService.GetServiceName(), &infraService)
	if err != nil {
		return nil, err
	}
	return infraService.BootService(ctx, input)
}

// UploadDeploymentPlan is the resolver for the uploadDeploymentPlan field.
func (r *mutationResolver) UploadDeploymentPlan(ctx context.Context, input *model.UploadDeploymentPlanInput) (*model.Infrastructure, error) {
	var infraService *service.InfrastructureService
	err := utils.ResolveService(infraService.GetServiceName(), &infraService)
	if err != nil {
		return nil, err
	}
	return infraService.UploadDeploymentPlan(ctx, input)
}

// Services is the resolver for the services field.
func (r *queryResolver) Services(ctx context.Context) ([]*model.Infrastructure, error) {
	panic(fmt.Errorf("not implemented"))
}
