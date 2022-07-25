package repository

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Key          string `gorm:"not null"`
	Description  string
	Technologies []Technology `gorm:"many2many:project_technologies;"`
}
