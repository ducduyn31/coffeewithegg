package sloky

import (
	"coffeewithegg/apps/adam/infrastructure/core/sloky/events"
)

type sloky struct {
	evManager *events.EventManager
}

func Sloky() *sloky {
	return &sloky{
		evManager: nil,
	}
}

func (self *sloky) WithCustomEventManager(evManager *events.EventManager) *sloky {
	self.evManager = evManager
	return self
}

func (self *sloky) GetEventManager() *events.EventManager {
	return self.evManager
}
