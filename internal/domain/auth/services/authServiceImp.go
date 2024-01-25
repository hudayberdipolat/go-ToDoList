package services

import (
	"errors"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/dto"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/repositories"
	"github.com/hudayberdipolat/go-ToDoList/internal/models"
	"github.com/hudayberdipolat/go-ToDoList/pkg/jwtToken"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authServiceImp struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return authServiceImp{
		authRepo: repo,
	}
}

func (a authServiceImp) Register(registerRequest dto.RegisterRequest) (*dto.AuthResponse, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 16)
	user := models.User{
		Username:  registerRequest.Username,
		FullName:  registerRequest.FullName,
		Password:  string(hashPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// register user
	if err := a.authRepo.CreateUser(user); err != nil {
		return nil, err
	}
	getUser, err := a.authRepo.GetUserWithUsername(registerRequest.Username)
	if err != nil {
		return nil, err
	}
	// generate token
	accessToken, errToken := jwtToken.GenerateToken(int(getUser.ID), getUser.Username)
	if errToken != nil {
		return nil, errToken
	}
	registerResponse := dto.NewAuthResponse(getUser, accessToken)
	return &registerResponse, nil
}

func (a authServiceImp) Login(loginRequest dto.LoginRequest) (*dto.AuthResponse, error) {
	getUser, err := a.authRepo.GetUserWithUsername(loginRequest.Username)
	if err != nil {
		return nil, errors.New("username ýa-da password nädogry!!!")
	}
	// Check password
	if errPassword := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(loginRequest.Password)); errPassword != nil {
		return nil, errors.New("username ýa-da password nädogry!!!")
	}
	// generate token
	accessToken, errToken := jwtToken.GenerateToken(int(getUser.ID), getUser.Username)
	if errToken != nil {
		return nil, errToken
	}
	loginResponse := dto.NewAuthResponse(getUser, accessToken)
	return &loginResponse, nil
}
