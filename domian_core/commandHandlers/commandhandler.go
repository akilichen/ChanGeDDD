package commandHandlers

import (
	"DDD/domian_core/commands"
)

type CommandHandler interface {
	Handle(command commands.Command) error
	// TODO 增加UnitOfWork做工作单元的执行，参照gormx包（stark集成）
}
