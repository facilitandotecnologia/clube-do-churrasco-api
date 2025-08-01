package ports

import (
	"api/internal/domain"
)

type UserRepository interface {
	Save(user *domain.User) error
	FindByID(id string) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}
