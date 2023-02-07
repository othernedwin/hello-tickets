package events

import (
	"time"
)

func NewEventTicketCreaetd(eventID string, ticketID string, ticketStatus string, ticketPrice float32, seatNumber string, metadata map[string]interface{}) *Event {
	return &Event{
		EventType: "TICKET_CREATED",
		Data: map[string]interface{}{
			"eventID":      eventID,
			"ticketID":     ticketID,
			"ticketStatus": ticketStatus,
			"ticketPrice":  ticketPrice,
			"seatNumber":   seatNumber,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
