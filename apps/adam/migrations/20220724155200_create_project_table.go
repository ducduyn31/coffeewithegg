package migrations

import (
	"coffeewithegg/apps/adam/project/repository"
	"gorm.io/gorm"
)

type Migration_20220724155200 struct {
}

func (m Migration_20220724155200) Name() string {
	return "20220724155200_create_project_table"
}

func (m Migration_20220724155200) Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&repository.Technology{}, &repository.Project{})
	if err != nil {
		return err
	}

	return nil
}

func (m Migration_20220724155200) CheckCondition(db *gorm.DB) bool {
	return !db.Migrator().HasTable(&repository.Technology{}) && !db.Migrator().HasTable(&repository.Project{})
}
