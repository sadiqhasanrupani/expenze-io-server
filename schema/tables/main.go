package tables

import (
	"database/sql"
	"log"

)

func CreateTables(DB *sql.DB) {
    createUsersTable := `
    CREATE TABLE IF NOT EXISTS users(
      id SERIAL PRIMARY KEY NOT NULL,
      first_name VARCHAR(100) NOT NULL,
      last_name VARCHAR(100) NOT NULL,
      email_id  VARCHAR(150) NOT NULL,
      password VARCHAR(255) NOT NULL
    )
  `
    _, err := DB.Exec(createUsersTable)
    if err != nil {
        log.Fatalf("Error creating users table: %v", err)
    }

  log.Println("tables created successfully")
}

