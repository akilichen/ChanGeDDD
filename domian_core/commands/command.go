package commands

type Command interface {
	IsValid() error // 验证系统数据
}
