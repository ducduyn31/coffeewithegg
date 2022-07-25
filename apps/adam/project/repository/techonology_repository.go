package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Technology struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string
	Description string
	Key         string     `gorm:"not null"`
	Projects    []*Project `gorm:"many2many:project_technologies;"`
}
