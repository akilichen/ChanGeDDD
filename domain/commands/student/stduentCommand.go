package student

import (
	"time"
)

type StdCommandBase struct {
	Id        string
	Name      string
	Email     string
	BirthDate time.Time
	Phone     string
	Province  string
	City      string
	Country   string
	Street    string
}
