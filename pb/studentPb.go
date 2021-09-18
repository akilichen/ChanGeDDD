package pb

import "time"

type StudentInterface interface {
	GetStudentById(string) (*GetStudentResponse, error)
	AddStudent(*AddStudentRequest) (*CommonResponse, error)
}

type CommonResponse struct {
}

type AddressItem struct {
	Province string
	City     string
	County   string
	Street   string
}

type GetStudentResponse struct {
	Id      string
	Name    string
	Age     int
	Address AddressItem
}

type AddStudentRequest struct {
	Name      string
	Email     string
	Phone     string
	BirthDate time.Time
	Address   AddressItem
}
