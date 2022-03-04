package storage

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEventAlreadyExists = errors.New("event already exists")
	ErrEventDoesNotExists = errors.New("event does not exist")
)

type Event struct {
	ID          uuid.UUID
	Title       string
	StartedAt   time.Time
	FinishedAt  time.Duration
	Description string
	UserID      uuid.UUID
	Notify      time.Duration
}

func NewEvent(
	title string,
	startedAt time.Time,
	finishedAt time.Duration,
	description string,
	userID uuid.UUID,
	notify time.Duration) *Event {
	return &Event{
		ID:          uuid.New(),
		Title:       title,
		StartedAt:   startedAt,
		FinishedAt:  finishedAt,
		Description: description,
		UserID:      userID,
		Notify:      notify,
	}
}
