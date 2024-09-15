package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"expenze-io.com/internal/repositories"
	"expenze-io.com/internal/services"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	ConnStr := os.Getenv("PG_CONNSTR")

	DB, err = sql.Open("postgres", ConnStr)

	if err != nil {
		log.Panic("Something went wrong in database: ", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Initializing repositories
	userRepo := repositories.NewUserRepository(DB)
	otpRepo := repositories.NewOtpRespository(DB)
	countryRepo := repositories.NewCountryRespository(DB)

	// Initializing services
	dbServices := services.NewDatabaseService(services.DatabaseService{
		UserRepo:    userRepo,
		OtpRepo:     otpRepo,
		CountryRepo: countryRepo,
	})

	err = dbServices.SetupDatabase()

	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	fmt.Println("Database is up and running")
}
