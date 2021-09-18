package modles

import (
	"DDD/domian_core/models"
)

type AddressVo struct {
	models.Vo
	Province string
	City     string
	County   string
	Street   string
}

func  NewAddress(province, city, country, street string) AddressVo {
	return AddressVo{
		Vo:       models.Vo{},
		Province: province,
		City:     city,
		County:   country,
		Street:   street,
	}
}
