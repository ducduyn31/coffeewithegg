package events

type EventManager interface {
	Publish(message EventMessage) error
	Subscribe(EventMessage, func(message EventMessage) error) error
}

type EventMessage interface {
}

type PollableEventManager interface {
	PollEvents() ([]EventMessage, error)
}

type WorkerPoolManager interface {
}

type Worker struct {
}
