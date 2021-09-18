package services

import "DDD/application/model"

type StdServices interface {
	GetStudentById(id string) (*model.GetStudentResponseDto, error)
	AddStudent(req *model.AddStudentRequestDto) error
}
