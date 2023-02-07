package events

import (
	"time"
)

func NewEventEventUpdated(eventID string, updateType string, updateData map[string]interface{}, metadata map[string]interface{}) *Event {
	return &Event{
		EventType: "EVENT_UPDATED",
		Data: map[string]interface{}{
			"eventID":    eventID,
			"updateType": updateType,
			"updateData": updateData,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
