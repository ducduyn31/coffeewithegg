package service

import (
	infraService "coffeewithegg/apps/adam/infrastructure/service"
	projectService "coffeewithegg/apps/adam/project/service"
	"coffeewithegg/apps/adam/utils"
)

type ServiceMap struct {
	services map[string]*utils.ServiceModule
}

func (serviceMap *ServiceMap) registerService(service utils.ServiceModule) {
	serviceMap.services[service.GetServiceName()] = &service
	utils.RegisterNewService(service)
}

var serviceMapInstance = &ServiceMap{
	services: make(map[string]*utils.ServiceModule),
}

func GetServiceMap() *map[string]*utils.ServiceModule {
	return &serviceMapInstance.services
}

func InitServicesContainer() {

	// Project Modules
	serviceMapInstance.registerService(&projectService.ProjectService{})
	serviceMapInstance.registerService(&projectService.TechnologyService{})

	// Infra modules
	serviceMapInstance.registerService(&infraService.InfrastructureService{})
	serviceMapInstance.registerService(&infraService.InfrastructureEventManager{})
}
