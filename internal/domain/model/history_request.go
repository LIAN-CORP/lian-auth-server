package model

import (
	"time"

	"github.com/google/uuid"
)

type HistoryRequest struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Status    string
	CreatedAt time.Time
	Reason    string
}
