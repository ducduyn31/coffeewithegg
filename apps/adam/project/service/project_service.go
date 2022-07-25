package service

import (
	"coffeewithegg/apps/adam/graph/model"
	"coffeewithegg/apps/adam/project/repository"
	"context"
	"github.com/google/wire"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"strconv"
)

type projectService struct {
	db *gorm.DB
}

func (service *projectService) GetProjects(ctx context.Context, filters *model.ProjectFilter) ([]*model.Project, error) {
	// Offset Filter
	var offset int
	if filters.Offset == nil || *filters.Offset <= 0 {
		offset = 0
	} else {
		offset = *filters.Offset
	}

	// Limit Filter
	var limit int
	if filters.Count == nil || *filters.Count <= 0 {
		limit = 25 // Default to 25 items per query
	} else {
		limit = *filters.Count
	}

	var projectDBOs []*repository.Project

	service.db.Find(&projectDBOs).Offset(offset).Limit(limit)

	return funk.Map(projectDBOs, mapProjectDBOtoProject).([]*model.Project), nil
}

func mapProjectDBOtoProject(projectDBO *repository.Project) *model.Project {
	return &model.Project{
		ID:           strconv.Itoa(int(projectDBO.ID)),
		Name:         projectDBO.Name,
		Description:  nil,
		Technologies: nil,
	}
}

func newProjectService(db *gorm.DB) *projectService {
	return &projectService{
		db: db,
	}
}

func NewProjectService() *projectService {
	wire.Build(newProjectService)
	return &projectService{}
}
