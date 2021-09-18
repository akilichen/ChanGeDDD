package studentRepo

import (
	"DDD/domain/modles"
	"log"
)

type StdRepo struct {
}

func NewStdRepo() *StdRepo {
	p := new(StdRepo)
	return p
}

func (p StdRepo) Create(obj interface{}) error {
	log.Printf("%+v", obj)
	return nil
}

func (p StdRepo) GetById(guid string) interface{} {
	return nil
}

func (p StdRepo) GetAll() []interface{} {
	return nil
}

func (p StdRepo) Update(obj interface{}) error {
	return nil
}

func (p StdRepo) Remove(guid string) error {
	return nil
}

func (p StdRepo) GetByEmail(email string) (*modles.StudentEntity, error) {
	return nil, nil
}
