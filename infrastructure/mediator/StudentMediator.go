package mediator

import (
	cmdHandlersImpl "DDD/domain/commandHandlers"
	eventHandlersImpl "DDD/domain/eventHandlers"
	"DDD/domain/events/domainNotification"
	"DDD/domian_core/commandHandlers"
	"DDD/domian_core/commands"
	"DDD/domian_core/eventHandlers"
	"DDD/domian_core/events"
	"DDD/helper/jsonHelper"
	"DDD/infrastructure/eventStore"
	"DDD/infrastructure/repository/studentRepo"
)

type StdMediator struct {
	cmdHandler   commandHandlers.CommandHandler
	eventHandler eventHandlers.EventHandler
}

func NewStdMediator() *StdMediator {
	p := new(StdMediator)
	p.cmdHandler = cmdHandlersImpl.NewStdCmdHandler(studentRepo.NewStdRepo(), p)
	p.eventHandler = eventHandlersImpl.NewStdEventHandler(studentRepo.NewStdRepo())
	return p
}

// SendCommand 可以直接进程内消费command，也可以将command发送到消息中间件实现
// 此处使用facade模式为命令提供进程内消费
// todo 读写命令分离，CQRS完成进程外数据消费（高并发系统考虑）(需要考虑怎么去设计)
func (p StdMediator) SendCommand(commands commands.Command) error {
	// 不做依赖注入，直接实例化，解耦开来，后续有需要可以改成队列消费模式
	err := p.cmdHandler.Handle(commands)
	if err != nil {
		return err
	}
	return nil
}

func (p StdMediator) RaiseEvents(events events.Event, eventType string) error {
	// 判断是否领域事件通知，不是则保存事件做事件回溯
	_, ok := events.(*domainNotification.DomainNotification)
	if !ok {
		err := eventStore.NewEventStore("chenhj", jsonHelper.MarshalToStringNoError(events), eventType).Save()
		if err != nil {
			return err
		}
	}

	err := p.eventHandler.Publish(events)
	if err != nil {
		return err
	}
	return nil
}
