package entity

import "github.com/google/uuid"

type Participant struct {
	ID       string
	Username string
}

func NewParticipant() Participant {
	return Participant{
		ID: uuid.NewString(),
	}
}
