package eventHandlers

import (
	studentEvent "DDD/domain/events/student"
	"DDD/domain/interfaces"
	event_core "DDD/domian_core/events"
	"log"
)

type StdEventHandler struct {
	stdRepo interfaces.IStdRepository
}

func NewStdEventHandler(stdRepo interfaces.IStdRepository) *StdEventHandler {
	p := new(StdEventHandler)
	p.stdRepo = stdRepo
	return p
}

func (p StdEventHandler) Publish(event event_core.Event) error {
	var err error
	switch event.(type) {
	case *studentEvent.RegisterStudentEvent:
		err = p.handleRegisterEvent(event.(*studentEvent.RegisterStudentEvent))
	default:

	}
	return err
}

func (p StdEventHandler) handleRegisterEvent(registerEvent *studentEvent.RegisterStudentEvent) error {

	// 注册成功
	log.Println("恭喜您，注册成功！")

	return nil
}

func (p StdEventHandler) handleUpdateEvent(updateEvent *studentEvent.RegisterStudentEvent) error {
	return nil
}

func (p StdEventHandler) handleRemoveEvent(removeEvent *studentEvent.RegisterStudentEvent) error {
	return nil
}
