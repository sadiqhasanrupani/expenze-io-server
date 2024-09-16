package models

import "time"

type Otp struct {
	ID        int       `json:"id"`
	OtpNumber int       `json:"otp_number"`
	ExpireAt  time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
