package events

import (
	"fmt"
	"time"

	"hello-tickets/src/utils"

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
		err = utils.UnmarshalAndValidate(eventData, &eventCancelledData)
		key = eventCancelledData.EventID
		event = NewEventCancelledEvent(key, metadata)
	case EventCreated:
		var eventCreatedData EventCreatedData
		err = utils.UnmarshalAndValidate(eventData, &eventCreatedData)
		key = "event-" + uuid.New().String()
		event = NewEventCreatedEvent(key, eventCreatedData.EventOrganizers, eventCreatedData.EventName, eventCreatedData.EventTime, eventCreatedData.EventLocation, metadata)
	case EventUpdated:
		var eventUpdatedData EventUpdatedData
		err = utils.UnmarshalAndValidate(eventData, &eventUpdatedData)
		key = eventUpdatedData.EventID
		event = NewEventUpdatedEvent(key, eventUpdatedData.UpdateTypes, eventUpdatedData.UpdateData, metadata)
	case TicketCreated:
		var ticketCreatedData TicketCreatedData
		err = utils.UnmarshalAndValidate(eventData, &ticketCreatedData)
		key = "ticket-" + uuid.New().String()
		event = NewTicketCreatedEvent(key, ticketCreatedData.EventID, ticketCreatedData.SeatNumber, ticketCreatedData.TicketStatus, ticketCreatedData.TicketPrice, metadata)
	case TicketUpdated:
		var ticketUpdatedData TicketUpdatedData
		err = utils.UnmarshalAndValidate(eventData, &ticketUpdatedData)
		key = ticketUpdatedData.TicketID
		event = NewTicketUpdatedEvent(key, ticketUpdatedData.UpdateTypes, ticketUpdatedData.UpdateData, metadata)
	case UserRegistered:
		var userRegisteredData UserRegisteredData
		err = utils.UnmarshalAndValidate(eventData, &userRegisteredData)
		key = "user-" + uuid.New().String()
		event = NewUserRegisteredEvent(key, userRegisteredData.FirstName, userRegisteredData.LastName, userRegisteredData.Email, userRegisteredData.PhoneNumber, metadata)
	default:
		err = fmt.Errorf("invalid event type %s", eventType)
	}

	return key, event, err
}
