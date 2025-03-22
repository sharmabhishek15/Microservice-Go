package services

import (
	"errors"
	"services/user-service/models"
	"services/user-service/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	existingUser, _ := s.Repo.GetUserByEmail(user.Email)
	if existingUser.Email != "" {
		return errors.New("user already exists")
	}

	return s.Repo.CreateUser(user)
}
