package events

import (
	"time"
)

func NewEventUserRegistered(userID string, firstName string, lastName string, email string, phoneNumber string, metadata map[string]interface{}) *Event {
	return &Event{
		EventType: "USER_REGISTERED",
		Data: map[string]interface{}{
			"userID":      userID,
			"firstName":   firstName,
			"lastName":    lastName,
			"email":       email,
			"phoneNumber": phoneNumber,
		},
		Timestamp: time.Now().UTC(),
		Metadata:  metadata,
	}
}