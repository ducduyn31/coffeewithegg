package service

import (
	projectService "coffeewithegg/apps/adam/project/service"
	"github.com/golobby/container/v3"
)

func InitServicesContainer() error {
	var err error

	// Project Modules
	err = container.Singleton(func() *projectService.TechnologyService {
		return projectService.NewTechnologyService()
	})
	if err != nil {
		return err
	}

	err = container.Singleton(func() *projectService.ProjectService {
		return projectService.NewProjectService()
	})
	if err != nil {
		return err
	}

	return nil
}
