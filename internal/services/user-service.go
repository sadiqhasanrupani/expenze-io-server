package services

import (
	"database/sql"
	"errors"

	"expenze-io.com/internal/body"
	"expenze-io.com/internal/models"
	"expenze-io.com/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userServiceRepo struct {
	userRepo    repositories.UserRepository
	countryRepo repositories.CountryRepo
}

type UserService struct {
	repo userServiceRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		repo: userServiceRepo{
			userRepo:    *repositories.NewUserRepository(db),
			countryRepo: *repositories.NewCountryRespository(db),
		},
	}
}

func (s *UserService) RegisterUser(req *body.RegistrationBody) error {
	existingUser, _ := s.repo.userRepo.FindByEmail(req.EmailID)

	if existingUser != nil {
		return errors.New("User already exists with this email")
	}

	if existingUser, _ := s.repo.userRepo.FindByMobileNum(req.MobilieNumber); existingUser != nil {
		return errors.New("User already exists with this mobilie number")
	}

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Find country id using given phone code
	country, err := s.repo.countryRepo.FindByPhoneCode(req.PhoneCode)

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
		CountryId:    country.ID,
		Validity:     false,
	}

	// Save the user in the database
	if err := s.repo.userRepo.Save(newUser); err != nil {
		return err
	}

	return nil
}

func (s *UserService) SendOtpMsg(body *body.RegistrationBody) (string, error) {
	existingUser, _ := s.repo.userRepo.FindByEmail(body.EmailID)

	if existingUser == nil {
		return "", errors.New("user not found which is related to this email")
	}

	maskedNumber := MaskPhoneNumber(existingUser.MobileNumber)

	return "We have send you a otp on your whatsapp mobile number which last four digit is " + maskedNumber, nil
}
