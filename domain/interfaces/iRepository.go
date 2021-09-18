package interfaces

type IRepository interface {
	Create(obj interface{}) error
	GetById(guid string) interface{}
	GetAll() []interface{}
	Update(obj interface{}) error
	Remove(guid string) error
}
