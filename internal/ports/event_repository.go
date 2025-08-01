package ports

import (
	"api/internal/domain"
)

type EventRepository interface {
	Save(event *domain.Event) error
	FindByID(id string) (*domain.Event, error)
	FindAll() ([]*domain.Event, error)
}
