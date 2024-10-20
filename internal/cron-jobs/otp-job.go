package cronjob

import "database/sql"

// create cron job for deleting expired otp
func DeleteExpiredOtp(db *sql.DB) error {
	query := `DELETE FROM otps WHERE expire_at < NOW()`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
