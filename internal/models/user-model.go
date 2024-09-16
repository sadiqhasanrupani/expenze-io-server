package models

import "time"

type User struct {
	ID           int       `json:"id"`
	FullName     string    `json:"full_name"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmailId      string    `json:"email_id"`
	MobileNumber string    `json:"mobile_number"`
	PhoneCode    string    `json:"phonecode"`
	CountryId    int       `json:"country_id"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	UpdateAt     time.Time `json:"updated_at"`
}

type MobileUser struct {
	ID int `json:"id"`
}
