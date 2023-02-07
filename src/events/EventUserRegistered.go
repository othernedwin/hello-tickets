package events

import (
	"time"
)

func NewEventUserRegistered(userID string, metadata map[string]interface{}) *Event {
	return &Event{
		EventType: "USER_REGISTERED",
		Data: map[string]interface{}{
			"userID": userID,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
