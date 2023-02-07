package events

import (
	"time"
)

func NewEventEventCreated(eventCreator string, eventID string, eventName string, eventTime string, eventLocation string, eventOrganizer string, tickets map[string]interface{}, metadata map[string]interface{}) *Event {
	return &Event{
		EventType: "EVENT_CREATED",
		Data: map[string]interface{}{
			"eventCreator":  eventCreator,
			"eventID":       eventID,
			"eventName":     eventName,
			"eventTime":     eventTime,
			"eventLocation": eventLocation,
			"tickets":       tickets,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
