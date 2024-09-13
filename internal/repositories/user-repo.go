package repositories

import (
	"database/sql"
	"log"

	"expenze-io.com/internal/models"
	"expenze-io.com/pkg"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// create table method
func (repo *UserRepository) CreateUserTable() error {
	createTableQuery := pkg.CreateTableQuery("users",
		` id SERIAL PRIMARY KEY NOT NULL,
      first_name VARCHAR(100) NOT NULL,
      last_name VARCHAR(100) NOT NULL,
      mobile_number VARCHAR(50) NOT NULL,
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

// FindByEmail searches for a user by email
func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User // Variable to hold the user

	// SQL query
	query := "SELECT id, CONCAT(first_name, ' ', last_name) AS full_name, email_id, password FROM users WHERE email_id = $1"

	// Execute query and get the row
	row := repo.db.QueryRow(query, email)

	// Scan the row into the user struct
	err := row.Scan(&user.ID, &user.FullName, &user.EmailId, &user.Password)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Save User
func (repo *UserRepository) Save(user *models.User) error {
	query := `INSERT INTO users (
    first_name,
    last_name,
    email_id,
    mobile_number,
    password
  ) VALUES ($1, $2, $3, $4, $5)`

	_, err := repo.db.Exec(query, user.FirstName, user.LastName, user.EmailId, user.MobileNumber, user.Password)

	if err != nil {
		return err
	}

	return nil
}