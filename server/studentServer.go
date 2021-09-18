package server

import (
	"DDD/application/model"
	"DDD/application/services"
	"DDD/application/services/impl"
	"DDD/pb"
)

type StudentServer struct {
	stdService services.StdServices
}

func NewStudentServer() *StudentServer {
	return &StudentServer{
		stdService: impl.NewStdService(),
	}
}

func (p *StudentServer) GetStudentById(id string) (resp *pb.GetStudentResponse, err error) {
	return nil, nil
}

func (p *StudentServer) AddStudent(req *pb.AddStudentRequest) (*pb.CommonResponse, error) {
	dto := &model.AddStudentRequestDto{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		BirthDate: req.BirthDate,
		Province:  req.Address.Province,
		City:      req.Address.City,
		County:    req.Address.County,
		Street:    req.Address.Street,
	}
	err := p.stdService.AddStudent(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CommonResponse{}, nil
}
