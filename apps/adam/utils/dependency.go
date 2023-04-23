package utils

import (
	"errors"
	"github.com/golobby/container/v3"
	"github.com/labstack/gommon/log"
	"reflect"
)

type ServiceModule interface {
	GetServiceName() string
	InitServiceInstance() error
	DeleteServiceInstance() error
}

func ResolveService(serviceName string, instance interface{}) error {
	var service ServiceModule
	err := container.NamedResolve(&service, serviceName)
	if err != nil {
		return err
	}

	receiverType := reflect.TypeOf(instance)
	if receiverType.Kind() != reflect.Ptr {
		return errors.New("instance must be a pointer")
	}
	reflect.ValueOf(instance).Elem().Set(reflect.ValueOf(service))

	return nil
}

func RegisterNewService(service ServiceModule) {
	err := container.NamedSingletonLazy(service.GetServiceName(), func() ServiceModule {
		log.Infof("Registering service: %s", service.GetServiceName())
		err := service.InitServiceInstance()
		if err != nil {
			log.Fatal("Failed to initialize service: ", err.Error())
			return nil
		}

		return service
	})
	if err != nil {
		log.Fatal(err)
	}
}
