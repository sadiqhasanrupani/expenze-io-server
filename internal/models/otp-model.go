package models

import "time"

type Otp struct {
	ID             int
	MobileOtp      int
	EmailOtp       int
	ExpireAt       time.Time
	EmailValidity  bool
	MobileValidity bool
	UserId         int64
	Token          string
	CreatedAt      time.Time
	UpdateAt       time.Time
}
