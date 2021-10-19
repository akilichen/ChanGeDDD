package domainNotification

type DomainNotification struct {
	Version int64
	Guid    string
	Key     string
	Value   string
}

func NewDomainNotification(key, value string) *DomainNotification {
	return &DomainNotification{
		Version: 1,
		Guid:    "snow id",
		Key:     key,
		Value:   value,
	}
}

func (p DomainNotification) PlaceHolder() {

}
