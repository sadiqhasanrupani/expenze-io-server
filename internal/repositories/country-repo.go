package repositories

import (
	"database/sql"
	"errors"
	"log"

	"expenze-io.com/internal/models"
	"github.com/biter777/countries"
	"github.com/lib/pq"
)

type CountryRepo struct {
	db *sql.DB
}

func NewCountryRespository(db *sql.DB) *CountryRepo {
	return &CountryRepo{db: db}
}

func (repo *CountryRepo) CreateCountryTable() error {
	query := `CREATE TABLE IF NOT EXISTS countries (
    id SERIAL PRIMARY KEY NOT NULL,
    iso CHAR(2) NOT NULL UNIQUE,
    name VARCHAR(80) NOT NULL,
    nicename VARCHAR(80) NOT NULL,
    iso3 CHAR(3),
    numcode SMALLINT,
    phonecode TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL 
  )`

	if _, err := repo.db.Exec(query); err != nil {
		log.Fatalf("Error creating countries table: %v", err)
		return err
	}

	return nil
}

func (repo *CountryRepo) InsertCountries() error {
	query := `SELECT COUNT(*) FROM countries`

	var count int
	err := repo.db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatalf("Failed to check country records: %v", err)
		return err
	}

	// If there are existing records, skip the insertion
	if count > 0 {
		log.Println("Countries already exist in the table. Skipping insertion.")
		return nil
	}

	query = `INSERT INTO countries (iso, name, nicename, iso3, numcode, phonecode)
          VALUES ($1, $2, $3, $4, $5, $6)
          ON CONFLICT (iso) DO NOTHING;`

	// Start transaction
	tx, err := repo.db.Begin()
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
		return err
	}

	// Loop through all countries from the countries package
	for _, country := range countries.All() {
		_, err = tx.Exec(query,
			country.Alpha2(),
			country.String(),
			country.String(),
			country.Alpha3(),
			int(country),
			pq.Array(country.CallCodes()))

		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Fatalf("Failed to rollback transaction: %v", rbErr)
			}
			log.Fatalf("Failed to insert country: %v", err)
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
		return err
	}

	log.Println("Successfully inserted all countries.")
	return nil
}

func (repo *CountryRepo) FindByPhoneCode(phonecode string) (*models.Country, error) {
	var country models.Country

	query := `
  SELECT id, name, iso, phonecode
  FROM countries
  WHERE $1 = ANY(phonecode)
  `

	row := repo.db.QueryRow(query, phonecode)

	err := row.Scan(&country.ID, &country.Name, &country.Iso, pq.Array(&country.PhoneCode))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("No rows are present related to the current phonecode")
		}

		return nil, err
	}

	return &country, nil

}
