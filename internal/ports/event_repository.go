package ports

import (
	"api/internal/domain"
	"context"
)

type EventRepository interface {
	Save(ctx context.Context, event *domain.Event) error
	FindByID(ctx context.Context, id string) (*domain.Event, error)
	FindAll(ctx context.Context) ([]*domain.Event, error)
}
