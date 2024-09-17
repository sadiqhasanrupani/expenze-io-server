package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"expenze-io.com/internal/body"
	"expenze-io.com/internal/mails"
	"expenze-io.com/internal/models"
	"expenze-io.com/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type authServiceRepo struct {
	userRepo    repositories.UserRepository
	countryRepo repositories.CountryRepo
	otpRepo     repositories.OtpRepository
}

type AuthService struct {
	repo authServiceRepo
}

func New(db *sql.DB) *AuthService {
	return &AuthService{
		repo: authServiceRepo{
			userRepo:    *repositories.NewUserRepository(db),
			countryRepo: *repositories.NewCountryRespository(db),
			otpRepo:     *repositories.NewOtpRespository(db),
		},
	}
}

func (as *AuthService) RegisterUser(req *body.RegistrationBody) (*int64, error) {
	existingUser, _ := as.repo.userRepo.FindByEmail(req.EmailID)

	if existingUser != nil {
		return nil, errors.New("User already exists with this email")
	}

	if existingUser, _ := as.repo.userRepo.FindByMobileNum(req.MobilieNumber); existingUser != nil {
		return nil, errors.New("User already exists with this mobilie number")
	}

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Find country id using given phone code
	country, err := as.repo.countryRepo.FindByPhoneCode(req.PhoneCode)

	if err != nil {
		return nil, err
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
	}

	// Save the user in the database
	userid, err := as.repo.userRepo.Save(newUser)

	if err != nil {
		log.Fatalln("user save error: ", err.Error())
		return nil, err
	}

	return userid, nil
}

func (s *AuthService) SendOtpMsg(body *body.RegistrationBody, userId int64) (string, error) {
	connStr := os.Getenv("PG_CONNSTR")
	companyName := os.Getenv("COMPANY_NAME")
	companyEmail := os.Getenv("COMPANY_EMAIL")

	waService, err := NewWhatsAppService(connStr)
	if err != nil {
		return "", err
	}

	phoneNumber := fmt.Sprintf("%s%s", body.PhoneCode, body.MobilieNumber)

	otpService := NewOTPService(6)
	sixDigitOtp, err := otpService.GenerateOTP()
	if err != nil {
		return "", err
	}

	otpCode, err := strconv.Atoi(sixDigitOtp)
	if err != nil {
		return "", err
	}

	// storing generated otp inside otp model
	newOtp := models.Otp{
		OtpNumber:      otpCode,
		ExpireAt:       time.Now().Add(10 * time.Minute),
		EmailValidity:  false,
		MobileValidity: false,
		UserId:         userId,
	}

	_, err = s.repo.otpRepo.New(&newOtp)
	if err != nil {
		return "", err
	}

	/*
	 ************************************
	 ********** SMTP EMAIL **************
	 ************************************
	 */

	// subject: "Your OTP Code for ExpenzeIO Verification",

	type TemplateData struct {
		OtpCode        string
		CompanyEmail   string
		ExpirationTime string
		Fullname       string
	}

	rooPath, err := os.Getwd()
	if err != nil {
		return "", nil
	}

	templatePath := filepath.Join(
		rooPath,
		"internal",
		"mails",
		"verify-otp",
		"index.html",
	)

	defaultChanlen := 2

	doneChans := make([]chan bool, defaultChanlen)
	errorChans := make([]chan error, defaultChanlen)

	for i, _ := range doneChans {
		doneChans[i] = make(chan bool)
	}

	for i, _ := range errorChans {
		errorChans[i] = make(chan error)
	}

	go mails.SendMailTemplate(&mails.SendMail{
		Subject:      "Your OTP Code for ExpenzeIO Verification",
		TemplatePath: templatePath,
		To:           []string{body.EmailID},
		TemplateData: TemplateData{
			OtpCode:        sixDigitOtp,
			CompanyEmail:   companyEmail,
			ExpirationTime: "10",
			Fullname:       body.Firstname + " " + body.Lastname,
		},
	}, doneChans[0], errorChans[0])

	/*
	 ************************************
	 ********** WHATSAPP OTP ************
	 ************************************
	 */
	message := fmt.Sprintf(`Dear %s %s,

Your OTP for completing the verification is *%v*. Please use this code within the next 10 minutes to proceed.

For your security, do not share this OTP with anyone.

Thank you,
%v
%v`, body.Firstname, body.Lastname, sixDigitOtp, companyName, companyEmail)

	go waService.SendMessage(phoneNumber, message, doneChans[1], errorChans[1])

	for i := 0; i < defaultChanlen; i++ {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				return "", err
			}

		case <-doneChans[i]:
			log.Printf("chan %v is done", i)
		}

	}

	maskedNumber := MaskPhoneNumber(body.MobilieNumber)
	maskedPhoneNum := fmt.Sprintf("+%s-%s", body.PhoneCode, maskedNumber)

	return fmt.Sprintf("We have send you a otp on your whatsapp mobile number which last four digit is %s", maskedPhoneNum), nil
}
