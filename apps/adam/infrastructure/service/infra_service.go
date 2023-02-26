package service

import (
	"coffeewithegg/apps/adam/graph/model"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golobby/container/v3"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"time"
)

type InfrastructureService struct {
	db  *gorm.DB
	rdb *redis.Client
}

func (service *InfrastructureService) BootService(ctx context.Context, input *model.RequestServiceInput) (*model.Infrastructure, error) {
	bootServiceKey := getServiceBootKey(input.Name)

	bootTx := func(tx *redis.Tx) error {
		keyExists := tx.Exists(ctx, bootServiceKey).Val()
		if keyExists == 1 {
			timeLeft := tx.TTL(ctx, bootServiceKey).Val()

			if timeLeft.Minutes() < 5 {
				extendedDuration := time.Hour + timeLeft
				tx.Expire(ctx, bootServiceKey, extendedDuration)
				log.Infof("Extending boot duration for service=%s to duration=%s", *input.Name, extendedDuration)
			}
		} else {
			msg := fmt.Sprintf("%d", time.Now().Unix())
			tx.Set(ctx, bootServiceKey, msg, time.Hour)
		}
		return nil
	}

	err := service.rdb.Watch(ctx, bootTx, bootServiceKey)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (service *InfrastructureService) UploadDeploymentPlan(ctx context.Context, input *model.UploadDeploymentPlanInput) (*model.Infrastructure, error) {
	return nil, nil
}

func getServiceBootKey(serviceName *string) string {
	return fmt.Sprintf("boot:%s", *serviceName)
}

func NewInfrastructureService() *InfrastructureService {
	var db *gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		log.Fatal("DB is not initialized")
		return nil
	}

	var rdb *redis.Client
	err = container.Resolve(&rdb)
	if err != nil {
		log.Fatal("Cache is not initialized")
		return nil
	}

	return &InfrastructureService{
		db:  db,
		rdb: rdb,
	}
}
