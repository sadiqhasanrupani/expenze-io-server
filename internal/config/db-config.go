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

	connStr := os.Getenv("PG_CONNSTR")

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Panic("Something went wrong in database: ", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Initailzing repositories
	userRepo := repositories.NewUserRepository(DB)
	otpRepo := repositories.NewOtpRespository(DB)

	// Initailzing services
	dbServices := services.NewDatabaseService(*userRepo, *otpRepo)

	err = dbServices.SetupDatabase()

	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	fmt.Println("Database is up and running")
}
