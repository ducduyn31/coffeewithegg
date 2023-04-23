package service

import (
	"coffeewithegg/apps/adam/shutdown"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"math/rand"
	"os"
	"time"
)

func (service *InfrastructureEventManager) startServiceWorker() {
	defer service.schedulerWaitGroup.Done()
	workerId, _ := uuid.NewUUID()
	hostName, _ := os.Hostname()
	name := fmt.Sprintf("ServiceWorker#%d-%s", 1+rand.Intn(viper.GetInt(`service-worker.count`)), hostName)
	for {
		select {
		case <-shutdown.MainContext().Done():
			log.Info("Shutdown signal received, stopping event worker ", name)
			return
		case <-time.After(time.Millisecond * time.Duration(viper.GetInt(`service-worker.interval`))):
			service.pollEvents(&workerId, name)
		}

	}
}

func (service *InfrastructureEventManager) workerPong(workerId *uuid.UUID) {
	timeout := viper.GetInt(`service-worker.timeout`)
	key := getCurrentTimeKey()
	service.rdb.SAdd(shutdown.MainContext(), key, workerId.String())
	service.rdb.Expire(shutdown.MainContext(), key, time.Millisecond*time.Duration(timeout))
}

func getCurrentTimeKey() string {
	timeout := viper.GetInt(`service-worker.timeout`)
	timeslot := time.Now().UnixMilli() / int64(timeout)
	return fmt.Sprintf("IE-%d", timeslot)
}

func (service *InfrastructureEventManager) pollEvents(workerId *uuid.UUID, workerName string) {
	service.workerPong(workerId)

	selectedWorker := service.rdb.SRandMember(shutdown.MainContext(), getCurrentTimeKey()).Val()
	if workerId.String() != selectedWorker {
		return
	}

	events := service.filterEventsByTime()
	if len(events) == 0 {
		return
	}

	service.assignEvents(events)

	log.Debug("Processing event by ", workerName, " events: ", events)
}

func (service *InfrastructureEventManager) assignEvents(events []EventMessage) {
	for _, event := range events {
		// Redis is single threaded, so this is safe
		successCount := service.rdb.ZRem(shutdown.MainContext(), service.evQueueName, event).Val()

		if successCount == 1 {
			service.handleEvent(event)
		}
	}
}

func (service *InfrastructureEventManager) filterEventsByTime() []EventMessage {
	topFive, _ := service.rdb.ZRangeWithScores(shutdown.MainContext(), service.evQueueName, 0, 5).Result()
	eventsMap, _ := convertToEventMessages(topFive)

	var result []EventMessage
	for event, timestamp := range eventsMap {
		if time.Now().After(time.Unix(0, timestamp)) {
			result = append(result, event)
		}
	}
	return result
}

func (service *InfrastructureEventManager) handleEvent(event EventMessage) {
	switch event.EventType {
	case BootInfrastructure:
		log.Info("Boot event received for service: ", event.ServiceName)
	}
}
