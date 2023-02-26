package service

import (
	infraService "coffeewithegg/apps/adam/infrastructure/service"
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

	// Infra modules
	err = container.Singleton(func() *infraService.InfrastructureService {
		return infraService.NewInfrastructureService()
	})
	if err != nil {
		return err
	}

	return nil
}
