package events

import (
	"time"
)

type UserRegisteredData struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func NewUserRegisteredEvent(userID string, firstName string, lastName string, email string, phoneNumber string, metadata map[string]interface{}) *Event {
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
