package events

import (
	"time"
)

type EventInterface interface {
	GetName() string
	GetTime() time.Time
	GetPayload() any
}

type EventHandlerInterface interface {
	Handler(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handle EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handle EventHandlerInterface) error
	Has(eventName string, handle EventHandlerInterface) bool
	Clear() error
}
