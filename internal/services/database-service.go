package services

import "expenze-io.com/internal/repositories"

type DatabaseService struct {
	UserRepo    *repositories.UserRepository
	OtpRepo     *repositories.OtpRepository
	CountryRepo *repositories.CountryRepo
}

func NewDatabaseService(props DatabaseService) *DatabaseService {
	return &DatabaseService{
		UserRepo:    props.UserRepo,
		OtpRepo:     props.OtpRepo,
		CountryRepo: props.CountryRepo,
	}
}

func (svc *DatabaseService) SetupDatabase() error {

	if err := svc.CountryRepo.CreateCountryTable(); err != nil {
		return err
	}

	if err := svc.UserRepo.CreateUserTable(); err != nil {
		return err
	}

	if err := svc.OtpRepo.CreateOtpTable(); err != nil {
		return err
	}

	if err := svc.CountryRepo.InsertCountries(); err != nil {
		return err
	}

	return nil
}
