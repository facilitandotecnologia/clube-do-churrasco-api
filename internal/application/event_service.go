package application

import (
	"api/internal/domain"
	"api/internal/ports"
	"context"
)

type EventService struct {
	Repo ports.EventRepository
}

func NewEventService(repo ports.EventRepository) *EventService {
	return &EventService{Repo: repo}
}

func (s *EventService) CreateEvent(event *domain.Event) error {
	return s.Repo.Save(context.Background(), event)
}

func (s *EventService) GetEvent(id string) (*domain.Event, error) {
	return s.Repo.FindByID(context.Background(), id)
}

func (s *EventService) ListEvents() ([]*domain.Event, error) {
	return s.Repo.FindAll(context.Background())
}
