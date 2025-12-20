package models

import (
	"encoding/json"
	"time"
)

const (
	notifOutMessageRelation = "notifications.out_messages"
)

// object model for insert/update
type NotifOutMessage struct {
	ID          int                        `json:"id" primaryKey:"true"`
	AppIS       int                        `json:"app_id"`
	Providers   []string                   `json:"providers"`
	Message     map[string]json.RawMessage `json:"message"`
	MessageType string                     `json:"message_type"`
	Status      int                        `json:"status"`
	Callback    bool                       `json:"callback"`
	CreatedAt   time.Time                  `json:"created_at"`
	Closed      bool                       `json:"closed"`
}

func (m NotifOutMessage) Relation() string {
	return notifOutMessageRelation
}

func (m NotifOutMessage) CollectionAgg() any {
	return &TotCount{0}
}

// object key model
type NotifOutMessageKey struct {
	Id    int `json:"id" required:"true"`
	AppId int `json:"app_id" required:"true"`
}

func (m NotifOutMessageKey) Relation() string {
	return notifOutMessageRelation
}
