package services

import "github.com/hudayberdipolat/go-ToDoList/internal/domain/user/dto"

type UserService interface {
	GetUserData(userID int) (*dto.UserResponse, error)
	UpdateUserData(userID int, updateData dto.UpdateUserData) error
	UpdateUserPassword(userID int, updatePassword dto.UpdateUserPassword) error
}
