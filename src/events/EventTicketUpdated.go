package events

import (
	"time"
)

func NewEventTicketUpdated(eventID string, ticketID string, updateType string, updateData map[string]interface{}, metadata map[string]interface{}) *Event {
	return &Event{
		EventType: "TICKET_UPDATED",
		Data: map[string]interface{}{
			"eventID":    eventID,
			"ticketID":   ticketID,
			"updateType": updateType,
			"updateData": updateData,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
