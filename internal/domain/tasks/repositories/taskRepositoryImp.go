package repositories

import (
	"errors"
	"github.com/hudayberdipolat/go-ToDoList/internal/models"
	"gorm.io/gorm"
)

type taskRepositoryImp struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return taskRepositoryImp{
		db: db,
	}
}

func (t taskRepositoryImp) GetAll(listID int) ([]models.Task, error) {
	var tasks []models.Task
	if err := t.db.Where("list_id = ?", listID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t taskRepositoryImp) GetOne(listID, taskID int) (*models.Task, error) {
	var task models.Task
	if err := t.db.Where("id = ?", taskID).Where("list_id =?", listID).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (t taskRepositoryImp) Create(task models.Task) error {
	if err := t.db.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func (t taskRepositoryImp) Update(listID, taskID int, task models.Task) error {
	var taskModel models.Task
	err := t.db.Model(&taskModel).Where("id =?", taskID).Where("list_id=?", listID).Updates(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (t taskRepositoryImp) Delete(listID, taskID int) error {
	getTask, err := t.GetOne(listID, taskID)
	if err != nil {
		return errors.New("Task-i öçürmekde näsazlyk döredi!!!")
	}
	if errDelete := t.db.Unscoped().Delete(&getTask).Error; err != nil {
		return errDelete
	}
	return nil
}
