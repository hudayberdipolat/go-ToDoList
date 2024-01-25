package repositories

import (
	"errors"
	"github.com/hudayberdipolat/go-ToDoList/internal/models"
	"gorm.io/gorm"
)

type userRepositoryImp struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryImp{
		db: db,
	}
}

func (u userRepositoryImp) UpdateData(userID int, user models.User) error {
	var userModel models.User
	if err := u.db.Model(&userModel).Where("id=?", userID).Updates(&user).Error; err != nil {
		return errors.New("Bu username ady eýýäm ulanylýar!!!")
	}
	return nil
}

func (u userRepositoryImp) UpdatePassword(userID int, password string) error {
	var user models.User
	if err := u.db.Model(&user).Where("id = ?", userID).Updates(models.User{
		Password: password}).Error; err != nil {
		return err
	}
	return nil
}

func (u userRepositoryImp) GetUser(userID int) (*models.User, error) {
	var user models.User
	if err := u.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
