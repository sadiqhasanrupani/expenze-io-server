package repositories

import (
	"database/sql"
	"log"

	"expenze-io.com/lib"
)

type OtpRepository struct {
	db *sql.DB
}

func NewOtpRespository(db *sql.DB) *OtpRepository {
	return &OtpRepository{db: db}
}

// method
func (repo *OtpRepository) CreateOtpTable() error {
	createOtpQuery := lib.CreateTableQuery("otp", `
    id SERIAL PRIMARY KEY NOT NULL,
    otp_number INTEGER UNIQUE NOT NULL,
  `)

	_, err := repo.db.Exec(createOtpQuery)

	if err != nil {
		log.Fatalf("Error creating OTP table: %v", err)
		return err
	}

	log.Println("OTP table created successfully")
	return nil
}