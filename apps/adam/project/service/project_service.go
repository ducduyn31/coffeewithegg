package service

import (
	"coffeewithegg/apps/adam/graph/model"
	"coffeewithegg/apps/adam/project/repository"
	"context"
	"errors"
	"fmt"
	"github.com/golobby/container/v3"
	"github.com/labstack/gommon/log"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProjectService struct {
	db          *gorm.DB
	techService *TechnologyService
}

func (service *ProjectService) GetProjects(_ context.Context, filters *model.ProjectFilter) ([]*model.Project, error) {
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

	service.db.Preload("Technologies").Offset(offset).Limit(limit).Find(&projectDBOs)

	return funk.Map(projectDBOs, mapProjectDBOtoProject).([]*model.Project), nil
}

func (service *ProjectService) UpsertProject(_ context.Context, input *model.ProjectInput) (*model.Project, error) {
	var err error

	err = validateProjectInput(input)
	if err != nil {
		return nil, err
	}

	// Get existing project
	var project *repository.Project
	if input.ID != nil {
		result := service.db.First(&project, input.ID)
		if result.RowsAffected == 0 {
			return nil, errors.New(fmt.Sprintf("Project Id=%s does not exist", *input.ID))
		}
	}

	// Update/Create the project
	err = service.db.Transaction(func(tx *gorm.DB) error {
		if project == nil {
			project = &repository.Project{}
		}

		if input.Key != nil {
			project.Key = *input.Key
		}

		if input.Name != nil {
			project.Name = *input.Name
		}

		if input.Description != nil {
			project.Description = *input.Description
		}

		if len(input.Technologies) > 0 {
			technologies, err := service.techService.UpsertTechnologiesTransaction(tx, input.Technologies)
			if err != nil {
				return err
			}
			project.Technologies = technologies
		}

		result := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&project)

		if result.RowsAffected == 0 {
			return result.Statement.Error
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return mapProjectDBOtoProject(project), nil
}

func NewProjectService() *ProjectService {
	var db *gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		log.Fatal("DB is not initialized")
		return nil
	}

	var techService *TechnologyService
	err = container.Resolve(&techService)
	if err != nil {
		log.Fatal("TechnologyService is not initialized")
		return nil
	}

	return &ProjectService{
		db:          db,
		techService: techService,
	}
}
