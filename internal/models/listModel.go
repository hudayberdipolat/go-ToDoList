package models

import "time"

type listStatus int

type List struct {
	ID         uint      `json:"id"`
	ListName   string    `json:"list_name"`
	ListStatus string    `json:"list_status"`
	UserID     uint      `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (*List) TableName() string {
	return "lists"
}
