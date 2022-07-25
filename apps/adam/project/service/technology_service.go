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

func (service *TechnologyService) UpsertTechnologiesTransaction(tx *gorm.DB, input []*model.TechnologyInput) ([]repository.Technology, error) {
	technologies := funk.Map(input, mapTechnologyToTechnologyDBO).([]*repository.Technology)

	tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(technologies, 200)

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
