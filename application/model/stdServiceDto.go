package model

import "time"

type GetStudentResponseDto struct {
}

type AddStudentRequestDto struct {
	Name      string
	Email     string
	Phone     string
	BirthDate time.Time
	Province  string
	City      string
	County    string
	Street    string
}
