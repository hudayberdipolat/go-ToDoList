package models

import "time"

type Task struct {
	ID         uint      `json:"id"`
	TaskName   string    `json:"task_name"`
	TaskStatus string    `json:"task_status"`
	ListID     uint      `json:"list_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (*Task) TableName() string {
	return "tasks"
}
