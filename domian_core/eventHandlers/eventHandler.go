package eventHandlers

import "DDD/domian_core/events"

type EventHandler interface {
	Publish(event events.Event) error
}
