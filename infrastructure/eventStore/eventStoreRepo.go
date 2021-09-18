package eventStore

import (
	"time"
)

type EventStore struct {
	id        string
	userName  string
	data      string
	dataType  string
	timestamp time.Time
}

func NewEventStore(userName, data, dataType string) *EventStore {
	return &EventStore{
		id:        "eventStoreSnowId",
		userName:  userName,
		data:      data,
		dataType:  dataType,
		timestamp: time.Now(),
	}
}

func (p EventStore) Save() error {
	// 根据实例化的p接受者完成事件数据落地入库
	return nil
}

func (p EventStore) GetEventDataById() (err error) {
	// 根据p接受者的id将数据库中数据获取并回填p接受者
	return nil
}
