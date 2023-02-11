package events

import (
	"time"
)

const (
	EventCancelled string = "EVENT_CANCELLED"
)

type EventCancelledData struct {
	EventID string `json:"eventID"`
}

func NewEventCancelledEvent(eventID string, metadata map[string]interface{}) Event {
	return Event{
		EventType: EventCancelled,
		Data: map[string]interface{}{
			"eventID": eventID,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
