package application

import (
	"api/internal/domain"
	"api/internal/ports"
)

type EventService struct {
	Repo ports.EventRepository
}

func NewEventService(repo ports.EventRepository) *EventService {
	return &EventService{Repo: repo}
}

func (s *EventService) CreateEvent(event *domain.Event) error {
	return s.Repo.Save(event)
}

func (s *EventService) GetEvent(id string) (*domain.Event, error) {
	return s.Repo.FindByID(id)
}

func (s *EventService) ListEvents() ([]*domain.Event, error) {
	return s.Repo.FindAll()
}
