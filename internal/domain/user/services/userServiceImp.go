package services

import (
	"errors"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/user/dto"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/user/repositories"
	"github.com/hudayberdipolat/go-ToDoList/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userServiceImp struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (u userServiceImp) GetUserData(userID int) (*dto.UserResponse, error) {
	getUser, err := u.userRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	userResponse := dto.NewUserResponse(*getUser)
	return &userResponse, nil
}

func (u userServiceImp) UpdateUserData(userID int, updateData dto.UpdateUserData) error {
	updateUser := models.User{
		Username:  updateData.Username,
		FullName:  updateData.FullName,
		UpdatedAt: time.Now(),
	}
	if err := u.userRepo.UpdateData(userID, updateUser); err != nil {
		return err
	}
	return nil
}

func (u userServiceImp) UpdateUserPassword(userID int, updatePassword dto.UpdateUserPassword) error {
	// check old password
	getUser, err := u.userRepo.GetUser(userID)
	if err != nil {
		return err
	}
	errOldPassword := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(updatePassword.OldPassword))
	if errOldPassword != nil {
		return errors.New("Köne password nädogry!!!")
	}

	// hashPassword
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(updatePassword.Password), 16)

	if errPasswordUpdate := u.userRepo.UpdatePassword(int(getUser.ID), string(hashPassword)); errPasswordUpdate != nil {
		return errPasswordUpdate
	}
	return nil
}
