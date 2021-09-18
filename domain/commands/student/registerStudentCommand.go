package student

import (
	"errors"
	"time"
)

type RegStdCommand struct {
	*StdCommandBase
}

func NewRegStdCommand(name, email string, birthDate time.Time, phone, province, city, country, street string) *RegStdCommand {
	p := new(RegStdCommand)
	p.StdCommandBase = &StdCommandBase{
		Name:      name,
		Email:     email,
		BirthDate: birthDate,
		Phone:     phone,
		Province:  province,
		City:      city,
		Country:   country,
		Street:    street,
	}
	return p
}

func (p RegStdCommand) IsValid() error {
	if err := p.validPhone(); err != nil {
		return err
	}
	if err := p.validName(); err != nil {
		return err
	}
	return nil
}

func (p RegStdCommand) validPhone() error {
	if len(p.Phone) != 11 {
		return errors.New("手机号不合法")
	}
	return nil
}

func (p RegStdCommand) validName() error {
	if len(p.Name) == 0 {
		return errors.New("名称不能为空")
	}
	return nil
}
