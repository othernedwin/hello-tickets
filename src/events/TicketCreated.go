package events

import (
	"time"
)

const (
	TicketCreated string = "TICKET_CREATED"
)

type TicketCreatedData struct {
	EventID      string  `json:"eventID" validate:"required"`
	SeatNumber   string  `json:"seatNumber" validate:"required"`
	TicketStatus string  `json:"ticketStatus" validate:"required"`
	TicketPrice  float32 `json:"ticketPrice" validate:"required"`
}

func NewTicketCreatedEvent(ticketID string, eventID string, seatNumber string, ticketStatus string, ticketPrice float32, metadata map[string]interface{}) Event {
	return Event{
		EventType: TicketCreated,
		Data: map[string]interface{}{
			"ticketID":     ticketID,
			"eventID":      eventID,
			"seatNumber":   seatNumber,
			"ticketStatus": ticketStatus,
			"ticketPrice":  ticketPrice,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
