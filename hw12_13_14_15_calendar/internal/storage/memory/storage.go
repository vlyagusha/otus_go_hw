package memorystorage

import (
	"sync"

	"github.com/google/uuid"
	"github.com/vlyagusha/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	mu     sync.RWMutex
	events map[uuid.UUID]storage.Event
}

func (s *Storage) Create(e storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[e.ID]; ok {
		return storage.ErrEventAlreadyExists
	}

	s.events[e.ID] = e
	return nil
}

func (s *Storage) Update(e storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[e.ID] = e
	return nil
}

func (s *Storage) Delete(id uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[id]; !ok {
		return storage.ErrEventDoesNotExists
	}

	delete(s.events, id)
	return nil
}

func (s *Storage) FindAll() ([]storage.Event, error) {
	events := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		events = append(events, event)
	}
	return events, nil
}

func New() *Storage {
	return &Storage{
		events: make(map[uuid.UUID]storage.Event),
	}
}
