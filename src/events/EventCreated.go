package events

import (
	"time"
)

const (
	EventCreated string = "EVENT_CREATED"
)

type EventCreatedData struct {
	EventOrganizers []string `json:"eventOrganizers"`
	EventName       string   `json:"eventName"`
	EventTime       string   `json:"eventTime"`
	EventLocation   string   `json:"eventLocation"`
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
