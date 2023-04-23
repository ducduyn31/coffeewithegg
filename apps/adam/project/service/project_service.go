package service

import (
	"coffeewithegg/apps/adam/graph/model"
	"coffeewithegg/apps/adam/project/repository"
	"coffeewithegg/apps/adam/utils"
	"context"
	"errors"
	"fmt"
	"github.com/golobby/container/v3"
	"github.com/labstack/gommon/log"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type ProjectService struct {
	db          *gorm.DB
	techService *TechnologyService
}

func (service *ProjectService) GetServiceName() string {
	return "ProjectService"
}

func (service *ProjectService) InitServiceInstance() error {
	var db *gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		return errors.New(fmt.Sprintf("DB is not initialized: %s", err.Error()))
	}

	var techService *TechnologyService
	err = utils.ResolveService(techService.GetServiceName(), &techService)
	if err != nil {
		return errors.New(fmt.Sprintf("TechnologyService is not initialized: %s", err.Error()))
	}

	service.db = db
	service.techService = techService
	return nil
}

func (service *ProjectService) DeleteServiceInstance() error {
	log.Info("Removing ProjectService instance")
	return nil
}

func (service *ProjectService) GetProjects(_ context.Context, filters *model.ProjectFilter) ([]*model.Project, error) {
	if filters == nil {
		filters = &model.ProjectFilter{}
	}

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
			result := tx.Create(&project)

			if result.RowsAffected == 0 {
				return result.Error
			}
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
			_, err := service.techService.UpsertTechnologiesForProjectTransaction(tx, input.Technologies, project)
			if err != nil {
				return err
			}
		} else {
			result := tx.Save(&project)
			if result.RowsAffected == 0 {
				return result.Error
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return mapProjectDBOtoProject(project), nil
}
