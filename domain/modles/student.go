package modles

import (
	"DDD/domian_core/models"
	"time"
)

type StudentEntity struct {
	*models.Entity
	Name      string
	Email     string
	Phone     string
	BirthDate time.Time
	Address   AddressVo
}

func NewStudent(name, email, phone string, birthDate time.Time, address AddressVo) *StudentEntity {
	return &StudentEntity{
		Entity:    models.NewEntity(),
		Name:      name,
		Email:     email,
		Phone:     phone,
		BirthDate: birthDate,
		Address:   address,
	}
}
