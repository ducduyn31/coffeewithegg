package events

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RDBEventManager struct {
	rdb             *redis.Client
	ctx             context.Context
	eventQueueName  string
	workerCount     int
	actions         map[EventMessage][]func(message EventMessage) error
	pollEventsCount int
}

func (self *RDBEventManager) PollEvents() ([]EventMessage, error) {
	result, err := self.rdb.ZRangeWithScores(self.ctx, self.eventQueueName, 0, int64(self.pollEventsCount)).Result()
	if err != nil {
		return nil, err
	}
	var messages []EventMessage
	for _, item := range result {
		if item.Score < float64(time.Now().Unix()) {
			messages = append(messages, item.Member.(EventMessage))
		}
	}
	return messages, nil
}

func (self *RDBEventManager) InitDefault() error {
	self.eventQueueName = "event_queue"
	self.workerCount = 10
	self.ctx = context.Background()
	return nil
}

func (self *RDBEventManager) Publish(message EventMessage) error {
	return self.rdb.ZAdd(self.ctx, self.eventQueueName, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: message,
	}).Err()
}

func (self *RDBEventManager) Subscribe(event EventMessage, f func(message EventMessage) error) error {
	self.actions[event] = append(self.actions[event], f)
	return nil
}

func (self *RDBEventManager) onShutdown() {

}

func NewEventManager(rdb *redis.Client) *RDBEventManager {
	manager := &RDBEventManager{
		rdb:             rdb,
		actions:         make(map[EventMessage][]func(message EventMessage) error),
		pollEventsCount: 100,
	}
	manager.InitDefault()
	return manager
}
