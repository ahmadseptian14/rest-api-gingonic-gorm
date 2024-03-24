package models

import "time"

type User struct{
	ID *int `json:"id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Name *string `json:"name"`
	Email *string `json:"email"`
	Address *string `json:"address"`
	BornDate *time.Time `json:"born_date"`
}
