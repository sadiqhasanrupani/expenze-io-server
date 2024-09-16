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
      mobile_number VARCHAR(50) UNIQUE NOT NULL,
      email_id VARCHAR(150) UNIQUE NOT NULL,
      password VARCHAR(255) NOT NULL,
      country_id INTEGER NOT NULL,
      FOREIGN KEY (country_id) REFERENCES countries(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    `,
	)

	_, err := repo.db.Exec(createTableQuery)

	if err != nil {
		log.Fatalf("Error creating user table: %v", err)
		return err
	}

	return nil
}

// FindByEmail searches for a user by email
func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User // Variable to hold the user

	// SQL query
	query := "SELECT id, CONCAT(first_name, ' ', last_name) AS full_name, email_id, password, mobile_number FROM users WHERE email_id = $1"

	// Execute query and get the row
	row := repo.db.QueryRow(query, email)

	// Scan the row into the user struct
	err := row.Scan(&user.ID, &user.FullName, &user.EmailId, &user.Password, &user.MobileNumber)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) FindByMobileNum(phonenumber string) (*models.MobileUser, error) {
	var user models.MobileUser // Variable to hold the user

	// SQL query
	query := `SELECT id from users WHERE mobile_number = $1`
	// Execute query and get the row
	row := repo.db.QueryRow(query, phonenumber)

	// Scan the row into the user struct
	err := row.Scan(&user.ID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Save User
func (repo *UserRepository) Save(user *models.User) (*int64, error) {
	query := `INSERT INTO users (
    first_name,
    last_name,
    email_id,
    mobile_number,
    password,
    country_id
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING id`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var id int64
	err = stmt.QueryRow(
		user.FirstName,
		user.LastName,
		user.EmailId,
		user.MobileNumber,
		user.Password,
		user.CountryId,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
}
