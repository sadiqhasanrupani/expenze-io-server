package services

import (
	"database/sql"
	"errors"

	"expenze-io.com/internal/body"
	"expenze-io.com/internal/models"
	"expenze-io.com/internal/repositories"
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

func (s *UserService) RegisterUser(req *body.RegistrationBody) error {
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
		PhoneCode:    req.PhoneCode,
		Validity:     false,
	}

	// Save the user in the database
	if err := s.repo.Save(newUser); err != nil {
		return err
	}

	return nil
}

func (s *UserService) SendOtpMsg(body *body.RegistrationBody) (string, error) {
	existingUser, _ := s.repo.FindByEmail(body.EmailID)

	if existingUser == nil {
		return "", errors.New("user not found which is related to this email")
	}

	maskedNumber := MaskPhoneNumber(existingUser.MobileNumber)

	return "We have send you a otp on your whatsapp mobile number which last four digit is " + maskedNumber, nil
}
