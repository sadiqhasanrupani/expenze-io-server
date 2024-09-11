package repositories

import (
	"database/sql"
	"log"

	"expenze-io.com/lib"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// method
func (repo *UserRepository) CreateUserTable() error {
	createTableQuery := lib.CreateTableQuery("users",
		` id SERIAL PRIMARY KEY NOT NULL,
      first_name VARCHAR(100) NOT NULL,
      last_name VARCHAR(100) NOT NULL,
      email_id VARCHAR(150) UNIQUE NOT NULL,
      password VARCHAR(255) NOT NULL,
    `,
	)

	_, err := repo.db.Exec(createTableQuery)

	if err != nil {
		log.Fatalf("Error creating user table: %v", err)
		return err
	}

	log.Println("Users table created successfully.")
	return nil
}
