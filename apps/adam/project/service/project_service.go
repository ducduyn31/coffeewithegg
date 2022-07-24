package service

import (
	"coffeewithegg/apps/adam/graph/model"
	"context"
	"github.com/google/wire"
)

type projectService struct {
}

func (service *projectService) GetProjects(ctx context.Context, filters *model.ProjectFilter) ([]*model.Project, error) {

	return nil, nil
}

func newProjectService() *projectService {
	return &projectService{}
}

func NewProjectService() *projectService {
	wire.Build(newProjectService)
	return &projectService{}
}
