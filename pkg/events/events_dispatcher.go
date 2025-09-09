package events

import (
	"errors"
	"slices"
)

var ErrorAlreadyRegister = errors.New("this event handlder already was register")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok && slices.Contains(ed.handlers[eventName], handler) {
		return ErrorAlreadyRegister
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHandlerInterface)
	if len(ed.handlers) != 0 {
		return errors.New("can not clean handlers")
	}
	return nil
}


func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok && slices.Contains(ed.handlers[eventName], handler) {
		return true
	}
	return false
}
