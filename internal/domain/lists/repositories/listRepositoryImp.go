package repositories

import (
	"errors"
	"github.com/hudayberdipolat/go-ToDoList/internal/models"
	"gorm.io/gorm"
)

type listRepositoryImp struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) ListRepository {
	return listRepositoryImp{
		db: db,
	}
}

func (l listRepositoryImp) FindAll(userID int) ([]models.List, error) {
	var lists []models.List
	if err := l.db.Where("user_id = ?", userID).Find(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func (l listRepositoryImp) FindOne(userID, listID int) (*models.List, error) {
	var list models.List
	if err := l.db.Where("id =?", listID).Where("user_id=?", userID).First(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func (l listRepositoryImp) Create(list models.List) error {
	if err := l.db.Create(&list).Error; err != nil {
		return err
	}
	return nil
}

func (l listRepositoryImp) Update(userID, listID int, list models.List) error {
	var listModel models.List
	err := l.db.Model(&listModel).Where("id =?", listID).Where("user_id=?", userID).Updates(&list).Error
	if err != nil {
		return err
	}
	return nil
}

func (l listRepositoryImp) Delete(userID, listID int) error {
	getList, errGetList := l.FindOne(userID, listID)
	if errGetList != nil {
		return errors.New("Siz bu listi öçürip bilmeýärsiňiz!!!")
	}
	err := l.db.Unscoped().Delete(&getList).Error
	if err != nil {
		return err
	}
	return nil
}
