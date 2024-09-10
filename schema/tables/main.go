package tables

import (
	"database/sql"
	"fmt"
	"log"
)

func defaultString(tableStr string) string {
	result := fmt.Sprintf(`
    %v
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
   )
  `, tableStr)

	return result
}

func CreateTables(DB *sql.DB) {
	createUsersTable := defaultString(`
     CREATE TABLE IF NOT EXISTS users(
      id SERIAL PRIMARY KEY NOT NULL,
      first_name VARCHAR(100) NOT NULL,
      last_name VARCHAR(100) NOT NULL,
      email_id  VARCHAR(150) UNIQUE  NOT NULL,
      password VARCHAR(255) NOT NULL,`)

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	log.Println("Users table created successfully")

	createOTPstable := defaultString(`
    CREATE TABLE IF NOT EXISTS otps(
      id SERIAL PRIMARY KEY NOT NULL,
      otp_number INTEGER UNIQUE NOT NULL,
    `)

	_, err = DB.Exec(createOTPstable)

	if err != nil {
		log.Fatalf("Error creating OTP table: %v", err)
	}

	log.Println("Otps table created successfully.")
}
