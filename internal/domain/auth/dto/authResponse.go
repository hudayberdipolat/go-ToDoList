package dto

import "github.com/hudayberdipolat/go-ToDoList/internal/models"

type AuthResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	AccessToken string `json:"access_token"`
	CreatedAt   string `json:"created_at"`
}

func NewAuthResponse(user *models.User, accessToken string) AuthResponse {
	return AuthResponse{
		ID:          int(user.ID),
		Username:    user.Username,
		FullName:    user.FullName,
		CreatedAt:   user.CreatedAt.Format("01-02-2006"),
		AccessToken: accessToken,
	}
}
