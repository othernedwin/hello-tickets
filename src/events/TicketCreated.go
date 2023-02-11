package events

import (
	"time"
)

const (
	TicketCreated string = "TICKET_CREATED"
)

type TicketCreatedData struct {
	EventID      string  `json:"eventID"`
	SeatNumber   string  `json:"seatNumber"`
	TicketStatus string  `json:"ticketStatus"`
	TicketPrice  float32 `json:"ticketPrice"`
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
