package models

import "time"

type Country struct {
	ID        int       `json:"id"`
	Iso       string    `json:"iso"`
	Name      string    `json:"name"`
	Nicename  string    `json:"nicename"`
	Iso3      string    `json:"iso3"`
	Numcode   int       `json:"numcode"`
	PhoneCode []string  `json:"phonecode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
