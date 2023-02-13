package events

import (
	"time"
)

const (
	UserRegistered string = "USER_REGISTERED"
)

type UserRegisteredData struct {
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

func NewUserRegisteredEvent(userID string, firstName string, lastName string, email string, phoneNumber string, metadata map[string]interface{}) Event {

	return Event{
		EventType: UserRegistered,
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
