package service

import (
	"coffeewithegg/apps/adam/graph/model"
	"coffeewithegg/apps/adam/project/repository"
	"errors"
	"fmt"
	"github.com/golobby/container/v3"
	"github.com/labstack/gommon/log"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TechnologyService struct {
	db *gorm.DB
}

func (service *TechnologyService) GetServiceName() string {
	return "TechnologyService"
}

func (service *TechnologyService) InitServiceInstance() error {
	var db *gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		return errors.New(fmt.Sprintf("DB is not initialized: %s", err.Error()))
	}

	service.db = db
	return nil
}

func (service *TechnologyService) DeleteServiceInstance() error {
	log.Info("Removing TechnologyService instance")
	return nil
}

func (service *TechnologyService) UpsertTechnologiesForProjectTransaction(tx *gorm.DB, input []*model.TechnologyInput, project *repository.Project) ([]repository.Technology, error) {
	technologies := funk.Map(input, mapTechnologyToTechnologyDBO).([]*repository.Technology)

	err := tx.Session(&gorm.Session{FullSaveAssociations: true}).
		Model(&project).
		Clauses(clause.OnConflict{UpdateAll: true}).
		Association("Technologies").
		Replace(technologies)

	if err != nil {
		return nil, err
	}

	return funk.Map(technologies, copyTechnologyDBO).([]repository.Technology), nil
}
