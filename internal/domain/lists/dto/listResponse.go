package dto

import "github.com/hudayberdipolat/go-ToDoList/internal/models"

type ListResponse struct {
	ID         uint   `json:"id"`
	ListName   string `json:"list_name"`
	ListStatus string `json:"list_status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func NewListResponse(list models.List) ListResponse {
	return ListResponse{
		ID:         list.ID,
		ListName:   list.ListName,
		ListStatus: list.ListStatus,
		CreatedAt:  list.CreatedAt.Format("01-02-2006"),
		UpdatedAt:  list.UpdatedAt.Format("01-02-2006"),
	}
}
