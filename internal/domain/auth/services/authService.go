package services

import "github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/dto"

type AuthService interface {
	Register(registerRequest dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(loginRequest dto.LoginRequest) (*dto.AuthResponse, error)
}
