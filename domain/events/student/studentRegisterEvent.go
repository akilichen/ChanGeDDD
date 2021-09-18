package student

import "time"

type RegisterStudentEvent struct {
	Id        string
	Name      string
	Email     string
	Phone     string
	BirthDate time.Time
	timestamp time.Time
}

func NewRegisterStudentEvent(name, email, phone string, birthDate time.Time) *RegisterStudentEvent {
	return &RegisterStudentEvent{
		Id:        "snowId",
		Name:      name,
		Email:     email,
		Phone:     phone,
		BirthDate: birthDate,
		timestamp: time.Now(),
	}
}

func (p RegisterStudentEvent) PlaceHolder() {

}
