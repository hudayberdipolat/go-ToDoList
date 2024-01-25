package dto

import "github.com/hudayberdipolat/go-ToDoList/internal/models"

type UserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        int(user.ID),
		Username:  user.Username,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt.Format("01-02-2006"),
		UpdatedAt: user.UpdatedAt.Format("01-02-2006"),
	}
}
