package repository

import "gorm.io/gorm"

type Technology struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Key         string     `gorm:"not null"`
	Projects    []*Project `gorm:"many2many:project_technologies;"`
}
