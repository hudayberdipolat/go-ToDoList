package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (*User) TableName() string {
	return "users"
}
