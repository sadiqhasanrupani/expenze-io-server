package services

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"expenze-io.com/internal/body"
	"expenze-io.com/internal/models"
	"expenze-io.com/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userServiceRepo struct {
	userRepo    repositories.UserRepository
	countryRepo repositories.CountryRepo
	otpRepo     repositories.OtpRepository
}

type UserService struct {
	repo userServiceRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		repo: userServiceRepo{
			userRepo:    *repositories.NewUserRepository(db),
			countryRepo: *repositories.NewCountryRespository(db),
			otpRepo:     *repositories.NewOtpRespository(db),
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
	connStr := os.Getenv("PG_CONNSTR")

	waService, err := NewWhatsAppService(connStr)
	if err != nil {
		return "", err
	}

	phoneNumber := fmt.Sprintf("%s%s", body.PhoneCode, body.MobilieNumber)

	companyName := os.Getenv("COMPANY_NAME")
	companyEmail := os.Getenv("COMPANY_EMAIL")

	otpService := NewOTPService(6)
	sixDigitOtp, err := otpService.GenerateOTP()
	if err != nil {
		return "", nil
	}

	otpCode, err := strconv.Atoi(sixDigitOtp)
	if err != nil {
		return "", err
	}

	// storing generated otp inside otp model
	otpModel := models.Otp{
		OtpNumber: otpCode,
		ExpireAt:  time.Now().Add(10 * time.Minute),
	}

	_, err = s.repo.otpRepo.New(&otpModel)
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf(`Dear %s %s,

Your OTP for completing the verification is *%v*. Please use this code within the next 10 minutes to proceed.

For your security, do not share this OTP with anyone.

Thank you,
%v
%v`, body.Firstname, body.Lastname, sixDigitOtp, companyName, companyEmail)

	if err := waService.SendMessage(phoneNumber, message); err != nil {
		return "", err
	}

	// if err := waService.SendOtpButtonMessage(phoneNumber, "2FA Authentication", message, "", "COPY OTP", sixDigitOtp); err != nil {
	// 	return "", err
	// }

	maskedNumber := MaskPhoneNumber(body.MobilieNumber)
	maskedPhoneNum := fmt.Sprintf("+%s-%s", body.PhoneCode, maskedNumber)

	return fmt.Sprintf("We have send you a otp on your whatsapp mobile number which last four digit is %s", maskedPhoneNum), nil
}
