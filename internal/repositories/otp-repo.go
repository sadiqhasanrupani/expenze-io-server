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
    mobile_otp INTEGER UNIQUE NOT NULL,
    email_otp INTEGER UNIQUE NOT NULL,
    email_validity BOOLEAN NOT NULL,
    mobile_validity BOOLEAN NOT NULL,
    expire_at TIMESTAMP NOT NULL,
    user_id INTEGER NOT NULL,
    token VARCHAR(40) NOT NULL,
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
    mobile_otp,
    email_otp,
    expire_at,
    email_validity,
    mobile_validity,
    user_id,
    token
  ) VALUES ($1, $2, $3, $4, $5, $6, $7)
  RETURNING id
  `

	stmt, err := repo.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	var id int64
	err = stmt.QueryRow(
		otp.MobileOtp,
		otp.EmailOtp,
		otp.ExpireAt,
		otp.EmailValidity,
		otp.MobileValidity,
		otp.UserId,
		otp.Token,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return proto.Int64(id), nil
}

func (repo *OtpRepository) FindByUserId(userId int64) (*models.Otp, error) {
	otpDetails := &models.Otp{}

	query := `
  SELECT id, mobile_otp, email_otp, expire_at, email_validity, mobile_validity
  FROM otps
  WHERE user_id = $1
  `

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(userId)

	if err := row.Scan(otpDetails.ID,
		otpDetails.MobileOtp,
		otpDetails.EmailOtp,
		otpDetails.ExpireAt,
		otpDetails.EmailValidity,
		otpDetails.MobileValidity,
	); err != nil {
		return nil, err
	}

	return otpDetails, nil
}

func (repo *OtpRepository) ApproveMobileOtp(mobileOtp int64, userId int64) error {
  query := `UPDATE otps SET mobile_validity = true WHERE mobile_otp = $1 AND user_id = $2`
  _, err := repo.db.Exec(query, mobileOtp, userId)
  return err
}
