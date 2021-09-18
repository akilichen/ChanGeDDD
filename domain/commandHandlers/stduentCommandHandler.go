package commandHandlers

import (
	"DDD/domain/commands/student"
	"DDD/domain/events"
	studentImpl "DDD/domain/events/student"
	"DDD/domain/interfaces"
	"DDD/domain/modles"
	"DDD/domian_core/bus"
	"DDD/domian_core/commands"
	"errors"
)

type StdCommandHandler struct {
	stdRepo interfaces.IStdRepository
	bus     bus.MediatorHandler
}

func NewStdCmdHandler(stdRepo interfaces.IStdRepository, handler bus.MediatorHandler) *StdCommandHandler {
	p := new(StdCommandHandler)
	p.stdRepo = stdRepo
	p.bus = handler
	return p
}

func (p StdCommandHandler) Handle(command commands.Command) error {
	var err error
	switch command.(type) {
	case *student.RegStdCommand:
		err = p.handleRegister(command.(*student.RegStdCommand))
	default:

	}
	return err
}

func (p StdCommandHandler) handleRegister(registerCmd *student.RegStdCommand) error {

	address := modles.NewAddress(registerCmd.Province, registerCmd.City, registerCmd.Country, registerCmd.Street)
	std := modles.NewStudent(registerCmd.Name, registerCmd.Email, registerCmd.Phone, registerCmd.BirthDate, address)

	stds, err := p.stdRepo.GetByEmail(registerCmd.Email)
	if err != nil {
		return err
	}
	if stds != nil {
		// TODO 如果错误了，做事件通知
		return errors.New("该邮箱已被使用，请检查")
	}

	if err := p.stdRepo.Create(std); err != nil {
		return err
	}

	// 成功提交后，做消息通知
	stdEvent := studentImpl.NewRegisterStudentEvent(registerCmd.Name, registerCmd.Email, registerCmd.Phone, registerCmd.BirthDate)
	err = p.bus.RaiseEvents(stdEvent, events.StdRegEvent)
	if err != nil {
		return err
	}
	return nil
}

func (p StdCommandHandler) handleUpdate(updateCmd student.RegStdCommand) error {
	return nil
}

func (p StdCommandHandler) handleRemove(removeCmd student.RegStdCommand) error {
	return nil
}
