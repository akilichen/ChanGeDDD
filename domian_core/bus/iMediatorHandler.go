package bus

import (
	"DDD/domian_core/commands"
	"DDD/domian_core/events"
)

type MediatorHandler interface {
	SendCommand(command commands.Command) error
	RaiseEvents(event events.Event, eventType string) error
}
