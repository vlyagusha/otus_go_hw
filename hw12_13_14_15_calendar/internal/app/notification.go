package app

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	EventID uuid.UUID
	Title   string
	Dt      time.Time
	UserID  string
}
