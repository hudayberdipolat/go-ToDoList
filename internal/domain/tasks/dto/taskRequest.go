package dto

type CreateTaskRequest struct {
	TaskName string `json:"task_name"`
}

type UpdateTaskRequest struct {
	TaskName   string `json:"task_name"`
	TaskStatus string `json:"task_status"`
}
