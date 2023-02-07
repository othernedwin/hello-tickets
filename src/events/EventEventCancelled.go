package events

import (
	"time"
)

func NewEventEventCancelled(eventID string, metadata map[string]interface{}) *Event {
	return &Event{
		EventType: "EVENT_CANCELLED",
		Data: map[string]interface{}{
			"eventID": eventID,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
