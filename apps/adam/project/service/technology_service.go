package service

import (
	"coffeewithegg/apps/adam/graph/model"
	"coffeewithegg/apps/adam/project/repository"
	"github.com/golobby/container/v3"
	"github.com/labstack/gommon/log"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TechnologyService struct {
	db *gorm.DB
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

func NewTechnologyService() *TechnologyService {
	var db *gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		log.Fatal("DB is not initialized")
		return nil
	}
	return &TechnologyService{
		db: db,
	}
}
