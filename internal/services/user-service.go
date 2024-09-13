package services

import (
	"database/sql"
	"errors"

	"expenze-io.com/internal/models"
	"expenze-io.com/internal/repositories"
	"expenze-io.com/internal/validators"
	"golang.org/x/crypto/bcrypt"
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

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Creating a user
	newUser := &models.User{
		FirstName:    req.Firstname,
		LastName:     req.Lastname,
		Password:     string(hashedPassword),
		EmailId:      req.EmailID,
		MobileNumber: req.MobilieNumber,
	}

	// Save the user in the database
  if err := s.repo.Save(newUser); err != nil {
    return err
  }
    
  return nil
}
