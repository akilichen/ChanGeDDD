package interfaces

import (
	"DDD/domain/modles"
)

type IStdRepository interface {
	IRepository
	GetByEmail(email string) (*modles.StudentEntity, error)
}
