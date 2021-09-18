package models

import "time"

type Entity struct {
	Guid int64
}

func NewEntity() *Entity {
	p := new(Entity)
	p.Guid = time.Now().Unix() // TODO 雪花id

	return p
}
