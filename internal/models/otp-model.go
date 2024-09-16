package models

import "time"

type Otp struct {
	ID             int       `json:"id"`
	OtpNumber      int       `json:"otp_number"`
	ExpireAt       time.Time `json:"expire_at"`
	EmailValidity  bool      `json:"email_validity"`
	MobileValidity bool      `json:"mobile_validity"`
	UserId         int64     `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdateAt       time.Time `json:"updated_at"`
}
