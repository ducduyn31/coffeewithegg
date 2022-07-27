package service

import (
	"coffeewithegg/apps/adam/graph/model"
	"coffeewithegg/apps/adam/project/repository"
	"errors"
	"github.com/google/uuid"
	"github.com/thoas/go-funk"
	"strconv"
)

// Mapper
func mapProjectDBOtoProject(projectDBO *repository.Project) *model.Project {
	if projectDBO == nil {
		return nil
	}

	description := funk.GetOrElse(projectDBO.Description, "").(string)
	technologies := funk.Map(projectDBO.Technologies, mapTechnologyDBOtoTechnology).([]*model.Technology)

	return &model.Project{
		ID:           strconv.Itoa(int(projectDBO.ID)),
		Key:          projectDBO.Key,
		Name:         projectDBO.Name,
		Description:  &description,
		Technologies: technologies,
	}
}

func mapTechnologyToTechnologyDBO(technology *model.TechnologyInput) *repository.Technology {
	if technology == nil {
		return nil
	}

	var myUuid uuid.UUID
	var name = ""
	var description = ""
	var key = ""
	var err error

	if technology.ID != nil {
		myUuid, err = uuid.Parse(*technology.ID)
	}

	if err != nil {
		panic("invalid id for technology")
	}

	if technology.Name != nil {
		name = *technology.Name
	}

	if technology.Description != nil {
		description = *technology.Description
	}

	if technology.Key != nil {
		key = *technology.Key
	}

	return &repository.Technology{
		ID:          funk.GetOrElse(myUuid, uuid.New()).(uuid.UUID),
		Name:        name,
		Description: description,
		Key:         key,
	}
}

func mapTechnologyDBOtoTechnology(technologyDBO repository.Technology) *model.Technology {

	description := funk.GetOrElse(technologyDBO.Description, "").(string)

	return &model.Technology{
		ID:          technologyDBO.ID.String(),
		Key:         technologyDBO.Key,
		Name:        technologyDBO.Name,
		Description: &description,
	}
}

func copyTechnologyDBO(technology *repository.Technology) repository.Technology {
	return *technology
}

// Validation
func validateProjectInput(input *model.ProjectInput) error {
	if input == nil || (input.ID == nil && (input.Key == nil || input.Name == nil)) {
		return errors.New("invalid input")
	}
	return nil
}
