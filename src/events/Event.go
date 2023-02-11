package events

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	EventType string                 `json:"eventType"`
	Data      map[string]interface{} `json:"data"`
	Timestamp time.Time              `json:"timestamp"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

func NewEvent(eventType string, eventData []byte, metadata map[string]interface{}) (string, Event, error) {
	var key string
	var event Event
	var err error = nil

	switch eventType {
	case EventCancelled:
		var eventCancelledData EventCancelledData
		err = json.Unmarshal(eventData, &eventCancelledData)
		key = eventCancelledData.EventID
		event = NewEventCancelledEvent(key, metadata)
	case EventCreated:
		var eventCreatedData EventCreatedData
		err = json.Unmarshal(eventData, &eventCreatedData)
		key = "event-" + uuid.New().String()
		event = NewEventCreatedEvent(key, eventCreatedData.EventOrganizers, eventCreatedData.EventName, eventCreatedData.EventTime, eventCreatedData.EventLocation, metadata)
	case EventUpdated:
		var eventUpdatedData EventUpdatedData
		err = json.Unmarshal(eventData, &eventUpdatedData)
		key = eventUpdatedData.EventID
		event = NewEventUpdatedEvent(key, eventUpdatedData.UpdateTypes, eventUpdatedData.UpdateData, metadata)
	case TicketCreated:
		var ticketCreatedData TicketCreatedData
		err = json.Unmarshal(eventData, &ticketCreatedData)
		key = "ticket-" + uuid.New().String()
		event = NewTicketCreatedEvent(key, ticketCreatedData.EventID, ticketCreatedData.SeatNumber, ticketCreatedData.TicketStatus, ticketCreatedData.TicketPrice, metadata)
	case TicketUpdated:
		var ticketUpdatedData TicketUpdatedData
		err = json.Unmarshal(eventData, &ticketUpdatedData)
		key = ticketUpdatedData.TicketID
		event = NewTicketUpdatedEvent(key, ticketUpdatedData.UpdateTypes, ticketUpdatedData.UpdateData, metadata)
	case UserRegistered:
		var userRegisteredData UserRegisteredData
		err = json.Unmarshal(eventData, &userRegisteredData)
		key = "user-" + uuid.New().String()
		event = NewUserRegisteredEvent(key, userRegisteredData.FirstName, userRegisteredData.LastName, userRegisteredData.Email, userRegisteredData.PhoneNumber, metadata)
	default:
		err = fmt.Errorf("invalid event type %s", eventType)
	}

	return key, event, err
}
