package service

import (
	"coffeewithegg/apps/adam/infrastructure/core/sloky/events"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/golobby/container/v3"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"sync"
	"time"
)

type InfrastructureEventManager struct {
	rdb                *redis.Client
	core               events.EventManager
	evQueueName        string
	schedulerWaitGroup *sync.WaitGroup
}

func (service *InfrastructureEventManager) GetServiceName() string {
	return "InfrastructureEventManager"
}

func (service *InfrastructureEventManager) InitServiceInstance() error {
	var rdb *redis.Client
	err := container.Resolve(&rdb)
	if err != nil {
		return err
	}

	service.rdb = rdb
	service.evQueueName = "infra_event_queue"
	service.core = events.NewEventManager(rdb)
	service.schedulerWaitGroup = &sync.WaitGroup{}
	return nil
}

func (service *InfrastructureEventManager) DeleteServiceInstance() error {
	log.Info("Removing InfrastructureEventManager instance")
	if service.schedulerWaitGroup != nil {
		service.schedulerWaitGroup.Wait()
	}
	return nil
}

const (
	BootInfrastructure = "boot_infra"
)

// SendBootRequest /**
//  * SendBootRequest is a function to send a boot request for a service
//  * @param serviceKey
//  * @return error
//  */
func (service *InfrastructureEventManager) SendBootRequest(serviceKey string) error {
	log.Info("Sending boot request for service: ", serviceKey)

	err := service.rdb.ZAdd(service.rdb.Context(), service.evQueueName, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: service.generateEventMessage(BootInfrastructure, serviceKey),
	}).Err()
	if err != nil {
		return err
	}

	return service.core.Publish(service.generateEventMessage(BootInfrastructure, serviceKey))
}

func (service *InfrastructureEventManager) generateEventMessage(eventType string, serviceName string) EventMessage {
	return EventMessage{
		EventType:   eventType,
		ServiceName: serviceName,
		Id:          uuid.New(),
	}
}

// StartServiceWorkers /**
//  * StartServiceWorkers is a function to handle the event queue
//  * @param workersCount
//  * @return void
//  */
func (service *InfrastructureEventManager) StartServiceWorkers(workersCount int) {
	for i := 0; i < workersCount; i++ {
		service.schedulerWaitGroup.Add(1)
		go service.startServiceWorker()
	}
}

type EventMessage struct {
	Id          uuid.UUID `json:"id"`
	EventType   string    `json:"event_type"`
	ServiceName string    `json:"service_name"`
}

func (e EventMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

func convertToEventMessages(data []redis.Z) (map[EventMessage]int64, error) {
	var result = make(map[EventMessage]int64)
	for _, d := range data {
		var ev EventMessage
		err := json.Unmarshal([]byte(d.Member.(string)), &ev)
		if err != nil {
			return nil, err
		}
		result[ev] = int64(d.Score)
	}
	return result, nil
}
