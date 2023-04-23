package service

import (
	"coffeewithegg/apps/adam/shutdown"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

func (service *InfrastructureEventManager) startServiceWorker() {
	defer service.schedulerWaitGroup.Done()
	name := fmt.Sprintf("ServiceWorker#%d", 1+rand.Intn(viper.GetInt(`service-worker.count`)))
	for {
		select {
		case <-shutdown.MainContext().Done():
			log.Info("Shutdown signal received, stopping event worker ", name)
			return
		case <-time.After(time.Millisecond * time.Duration(viper.GetInt(`service-worker.interval`))):
			service.processEvent(name)
		}

	}
}

func (service *InfrastructureEventManager) processEvent(workerName string) {
	log.Debug("Processing event by ", workerName)
}
