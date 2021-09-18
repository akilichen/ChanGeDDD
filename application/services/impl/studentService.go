package impl

import (
	"DDD/application/model"
	"DDD/domain/commands/student"
	"DDD/domain/interfaces"
	"DDD/domian_core/bus"
	"DDD/infrastructure/mediator"
	"DDD/infrastructure/repository/studentRepo"
)

type StdService struct {
	bus  bus.MediatorHandler
	repo interfaces.IStdRepository
}

func NewStdService() *StdService {
	p := new(StdService)

	p.bus = mediator.NewStdMediator()
	p.repo = studentRepo.NewStdRepo()

	return p
}

func (p *StdService) GetStudentById(id string) (*model.GetStudentResponseDto, error) {
	p.repo.GetById(id)
	return nil, nil
}

func (p *StdService) AddStudent(req *model.AddStudentRequestDto) error {
	registerCmd := student.NewRegStdCommand(req.Name, req.Email, req.BirthDate, req.Phone, req.Phone, req.Province, req.City, req.Street)
	// 确保数据正确后，才进行命令发送
	err := registerCmd.IsValid()
	if err != nil {
		return err
	}
	// 发送命令
	err = p.bus.SendCommand(registerCmd)
	if err != nil {
		return err
	}
	return nil
}
