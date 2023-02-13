package events

import (
	"time"
)

const (
	EventCreated string = "EVENT_CREATED"
)

type EventCreatedData struct {
	EventOrganizers []string `json:"eventOrganizers" validate:"required"`
	EventName       string   `json:"eventName" validate:"required"`
	EventTime       string   `json:"eventTime" validate:"required"`
	EventLocation   string   `json:"eventLocation" validate:"required"`
}

func NewEventCreatedEvent(eventID string, eventOrganizers []string, eventName string, eventTime string, eventLocation string, metadata map[string]interface{}) Event {
	return Event{
		EventType: EventCreated,
		Data: map[string]interface{}{
			"eventID":         eventID,
			"eventOrganizers": eventOrganizers,
			"eventName":       eventName,
			"eventTime":       eventTime,
			"eventLocation":   eventLocation,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
