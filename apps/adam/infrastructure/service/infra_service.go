package service

import (
	"coffeewithegg/apps/adam/graph/model"
	"coffeewithegg/apps/adam/utils"
	"context"
	"errors"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/go-redis/redis/v8"
	"github.com/golobby/container/v3"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"time"
)

type InfrastructureService struct {
	db        *gorm.DB
	rdb       *redis.Client
	evManager *InfrastructureEventManager
}

func (service *InfrastructureService) GetServiceName() string {
	return "InfrastructureService"
}

func (service *InfrastructureService) InitServiceInstance() error {
	var db *gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		return errors.New(fmt.Sprintf("DB is not initialized: %s", err.Error()))
	}

	var rdb *redis.Client
	err = container.Resolve(&rdb)
	if err != nil {
		return errors.New(fmt.Sprintf("Redis is not initialized: %s", err.Error()))
	}

	var evManager *InfrastructureEventManager
	err = utils.ResolveService(evManager.GetServiceName(), &evManager)
	if err != nil {
		return errors.New(fmt.Sprintf("InfrastructureEventManager is not initialized: %s", err.Error()))
	}

	service.db = db
	service.rdb = rdb
	service.evManager = evManager
	return nil
}

func (service *InfrastructureService) DeleteServiceInstance() error {
	log.Info("Removing InfrastructureService instance")
	return nil
}

// BootService /**
//  * BootService is a function to boot a service
//  * @param ctx
//  * @param input
//  * @return *model.Infrastructure
//  * @return error
//  */
func (service *InfrastructureService) BootService(ctx context.Context, input *model.RequestServiceInput) (*model.Infrastructure, error) {
	bootServiceKey := getServiceBootKey(input.Name)

	bootTx := func(tx *redis.Tx) error {
		keyExists := tx.Exists(ctx, bootServiceKey).Val()
		if keyExists == 1 {
			return service.extendServiceUpTime(tx, ctx, bootServiceKey)
		}

		return service.bootAndTimeService(tx, ctx, bootServiceKey)
	}

	err := retry.Do(func() error {
		return service.rdb.Watch(ctx, bootTx, bootServiceKey)
	}, retry.Attempts(3), retry.Delay(1*time.Second))

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func getServiceBootKey(serviceName *string) string {
	return fmt.Sprintf("boot:%s", *serviceName)
}

func (service *InfrastructureService) extendServiceUpTime(tx *redis.Tx, ctx context.Context, serviceKey string) error {
	timeLeft := tx.TTL(ctx, serviceKey).Val()

	if timeLeft.Minutes() < 5 {
		extendedDuration := time.Hour + timeLeft
		_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Expire(ctx, serviceKey, extendedDuration)
			return nil
		})
		if err != nil {
			return err
		}

		log.Infof("Extending boot duration for service=%s to duration=%s", serviceKey, extendedDuration)
	}

	return nil
}

func (service *InfrastructureService) bootAndTimeService(tx *redis.Tx, ctx context.Context, serviceKey string) error {
	bootTime := fmt.Sprintf("%d", time.Now().Unix())
	upTime := time.Hour
	_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Set(ctx, serviceKey, bootTime, upTime)
		return nil
	})
	if err != nil {
		return err
	}
	log.Infof("Booting service=%s", serviceKey)

	err = service.evManager.SendBootRequest(serviceKey)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// UploadDeploymentPlan /**
//  * UploadDeploymentPlan is a function to upload a deployment plan
//  * @param ctx
//  * @param input
//  * @return *model.Infrastructure
//  * @return error
//  */
func (service *InfrastructureService) UploadDeploymentPlan(ctx context.Context, input *model.UploadDeploymentPlanInput) (*model.Infrastructure, error) {
	return nil, nil
}
