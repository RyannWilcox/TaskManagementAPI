package services

import (
	"task-mgmt/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type TaskService interface {
	CreateTask(db *gorm.DB, task models.Task) error
	UpdateTask(db *gorm.DB, taskID uuid.UUID, userID uuid.UUID) error
	DeleteTask(db *gorm.DB, taskID uuid.UUID, userID uuid.UUID) error
	GetTaskByID(db *gorm.DB, taskID uuid.UUID, userID uuid.UUID) (models.Task, error)
	GetTasksByUser(db *gorm.DB, userID uuid.UUID) error
	GetTasks(db *gorm.DB) error
}
type TaskServiceImpl struct {
}

func NewTaskService() *TaskServiceImpl {
	return &TaskServiceImpl{}
}

// CreateTask implements [TaskService].
func (t *TaskServiceImpl) CreateTask(db *gorm.DB, task models.Task) error {
	// Begin the transaction
	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(&task).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Attempt to commit to the db.
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// DeleteTask implements [TaskService].
func (t *TaskServiceImpl) DeleteTask(db *gorm.DB, taskID uuid.UUID, userID uuid.UUID) error {
	panic("unimplemented")
}

// GetTaskByID implements [TaskService].
func (t *TaskServiceImpl) GetTaskByID(db *gorm.DB, taskID uuid.UUID, userID uuid.UUID) (models.Task, error) {
	var task models.Task
	if err := db.Where("id = ? AND user_id = ?").First(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

// GetTasks implements [TaskService].
func (t *TaskServiceImpl) GetTasks(db *gorm.DB) error {
	panic("unimplemented")
}

// GetTasksByUser implements [TaskService].
func (t *TaskServiceImpl) GetTasksByUser(db *gorm.DB, userID uuid.UUID) error {
	panic("unimplemented")
}

// UpdateTask implements [TaskService].
func (t *TaskServiceImpl) UpdateTask(db *gorm.DB, taskID uuid.UUID, userID uuid.UUID) error {
	panic("unimplemented")
}
