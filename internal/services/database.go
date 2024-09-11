package services

import "expenze-io.com/internal/repositories"

type DatabaseService struct {
	userRepo *repositories.UserRepository
	otpRepo  *repositories.OtpRepository
}

func NewDatabaseService(userRepo repositories.UserRepository, otpRepo repositories.OtpRepository) *DatabaseService {
	return &DatabaseService{
		userRepo: &userRepo,
		otpRepo:  &otpRepo,
	}
}

func (svc *DatabaseService) SetupDatabase() error {
	err := svc.userRepo.CreateUserTable()

	if err != nil {
		return err
	}

	err = svc.otpRepo.CreateOtpTable()

	if err != nil {
		return err
	}

	return nil
}
