package repositories

import (
	"database/sql"
	"log"

	"expenze-io.com/internal/models"
	"expenze-io.com/pkg"
	"google.golang.org/protobuf/proto"
)

type OtpRepository struct {
	db *sql.DB
}

func NewOtpRespository(db *sql.DB) *OtpRepository {
	return &OtpRepository{db: db}
}

// method
func (repo *OtpRepository) CreateOtpTable() error {
	createOtpQuery := pkg.CreateTableQuery("otps", `
    id SERIAL PRIMARY KEY NOT NULL,
    otp_number INTEGER UNIQUE NOT NULL,
    email_validity BOOLEAN NOT NULL,
    mobile_validity BOOLEAN NOT NULL,
    expire_at TIMESTAMP NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
  `)

	_, err := repo.db.Exec(createOtpQuery)

	if err != nil {
		log.Fatalf("Error creating OTP table: %v", err)
		return err
	}

	return nil
}

func (repo *OtpRepository) New(otp *models.Otp) (*int64, error) {
	query := `
  INSERT INTO otps (
    otp_number,
    expire_at,
    email_validity,
    mobile_validity,
    user_id
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING id
  `

	stmt, err := repo.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	var id int64
	err = stmt.QueryRow(
		otp.OtpNumber,
		otp.ExpireAt,
		otp.EmailValidity,
		otp.MobileValidity,
		otp.UserId,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return proto.Int64(id), nil
}
