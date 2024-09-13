package services

import (
	"database/sql"
	"errors"

	"expenze-io.com/internal/repositories"
	"expenze-io.com/internal/validators"
)

type UserService struct {
  repo repositories.UserRepository
}

func NewUserService(db *sql.DB) *UserService {
  return &UserService{
    repo: *repositories.NewUserRepository(db),
  }
}

func (s *UserService) RegisterUser(req *validators.RegistrationBody) error {
  existingUser, _ := s.repo.FindByEmail(req.EmailID)

  if existingUser != nil {
    return errors.New("User already exists with the email")
  }

  return nil
}
