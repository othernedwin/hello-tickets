package events

import (
	"time"
)

const (
	TicketUpdated string = "TICKET_UPDATED"
)

const (
	Available string = "AVAILABLE"
	Locked    string = "LOCKED" // LOCKED-BY-USER-x-x-x
	Sold      string = "SOLD"   // SOLD-TO-USER-x-x-x
	Cancelled string = "CANCELLED"
)

const (
	UpdateTicketStatus string = "UPDATE_TICKET_STATUS"
	UpdateTicketPrice  string = "UPDATE_TICKET_PRICE"
)

type TicketUpdatedData struct {
	TicketID    string                 `json:"ticketID" validate:"required"`
	UpdateTypes []string               `json:"updateTypes" validate:"required"`
	UpdateData  map[string]interface{} `json:"updateData" validate:"required"`
}

func NewTicketUpdatedEvent(ticketID string, updateTypes []string, updateData map[string]interface{}, metadata map[string]interface{}) Event {
	return Event{
		EventType: TicketUpdated,
		Data: map[string]interface{}{
			"ticketID":    ticketID,
			"updateTypes": updateTypes,
			"updateData":  updateData,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
