package events

import (
	"time"
)

const (
	EventUpdated string = "EVENT_UPDATED"
)

const (
	UpdateEventOrganziers string = "UPDATE_EVENT_ORGANIZERS"
	UpdateEventName       string = "UPDATE_EVENT_NAME"
	UpdateEventTime       string = "UPDATE_EVENT_TIME"
	UpdateEventLocation   string = "UPDATE_EVENT_LOCATION"
)

type EventUpdatedData struct {
	EventID     string                 `json:"eventID" validate:"required"`
	UpdateTypes []string               `json:"updateTypes" validate:"required"`
	UpdateData  map[string]interface{} `json:"updateData" validate:"required"`
}

func NewEventUpdatedEvent(eventID string, updateTypes []string, updateData map[string]interface{}, metadata map[string]interface{}) Event {
	return Event{
		EventType: EventUpdated,
		Data: map[string]interface{}{
			"eventID":     eventID,
			"updateTypes": updateTypes,
			"updateData":  updateData,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
