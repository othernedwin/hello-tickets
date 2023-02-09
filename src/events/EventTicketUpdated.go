package events

import (
	"time"
)

const (
	Available string = "AVAILABLE"
	Locked    string = "LOCKED" // LOCKED-BY-USER-x-x-x
	Sold      string = "SOLD"   // SOLD-TO-USER-x-x-x
	CANCELLED string = "CANCELLED"
)

const (
	UpdateTicketStatus string = "UPDATE_TICKET_STATUS"
	UpdateTicketPrice  string = "UPDATE_TICKET_PRICE"
)

type TicketUpdatedData struct {
	UpdateTypes []string               `json:"updateTypes"`
	UpdateData  map[string]interface{} `json:"updateData"`
}

func NewTicketUpdatedEvent(ticketID string, eventID string, updateTypes []string, updateData map[string]interface{}, metadata map[string]interface{}) *Event {
	return &Event{
		EventType: "TICKET_UPDATED",
		Data: map[string]interface{}{
			"eventID":     eventID,
			"ticketID":    ticketID,
			"updateTypes": updateTypes,
			"updateData":  updateData,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}
