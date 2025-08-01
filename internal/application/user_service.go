package application

import (
    "api/internal/domain"
    "api/internal/ports"
)

type UserService struct {
    Repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
    return s.Repo.Save(user)
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
    return s.Repo.FindByID(id)
}

func (s *UserService) ListUsers() ([]*domain.User, error) {
    return s.Repo.FindAll()
}

func (s *UserService) FindByEmail(email string) (*domain.User, error) {
    return s.Repo.FindByEmail(email)
}