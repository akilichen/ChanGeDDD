package events

type EventStoreRepository interface {
	Save() error
	GetEventDataById() error
}
