package events

import (
	"time"
)

const (
	UpdateEventOrganziers string = "UPDATE_EVENT_ORGANIZERS"
	UpdateEventName       string = "UPDATE_EVENT_NAME"
	UpdateEventTime       string = "UPDATE_EVENT_TIME"
	UpdateEventLocation   string = "UPDATE_EVENT_LOCATION"
)

type EventUpdatedData struct {
	UpdateTypes []string               `json:"updateTypes"`
	UpdateData  map[string]interface{} `json:"updateData"`
}

func NewEventUpdatedEvent(eventID string, updateTypes []string, updateData map[string]interface{}, metadata map[string]interface{}) *Event {
	return &Event{
		EventType: "EVENT_UPDATED",
		Data: map[string]interface{}{
			"eventID":     eventID,
			"updateTypes": updateTypes,
			"updateData":  updateData,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
