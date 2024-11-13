// internal/application/user_service.go
package application

import (
	"github.com/treewalkr/gymtrack/internal/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *domain.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
	return s.repo.GetUserByID(id)
}
