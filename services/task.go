package services

import (
	"task-mgmt/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type TaskService interface {
	CreateTask(db *gorm.DB, task models.Task) error
	UpdateTask(db *gorm.DB, taskID uuid.UUID, userID uuid.UUID, update models.TaskUpdate) (models.Task, error)
	DeleteTask(db *gorm.DB, taskID uuid.UUID) error
	GetTaskByID(db *gorm.DB, taskID uuid.UUID, userID uuid.UUID) (models.Task, error)
	GetTasksByUser(db *gorm.DB, userID uuid.UUID) ([]models.Task, error)
	GetTasks(db *gorm.DB) ([]models.Task, error)
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
func (t *TaskServiceImpl) DeleteTask(db *gorm.DB, taskID uuid.UUID) error {
	var task models.Task

	if err := db.Where("id = ?").First(&task).Error; err != nil {
		return err
	}

	// begin the transaction
	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Delete(&task).Error; err != nil {
		tx.Rollback()
		return err
	}

	// commit the deletion
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// GetTaskByID implements [TaskService].
func (t *TaskServiceImpl) GetTaskByID(db *gorm.DB, taskID uuid.UUID, userID uuid.UUID) (models.Task, error) {
	var task models.Task
	if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

// GetTasks implements [TaskService].
func (t *TaskServiceImpl) GetTasks(db *gorm.DB) ([]models.Task, error) {
	var tasks []models.Task

	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTasksByUser implements [TaskService].
func (t *TaskServiceImpl) GetTasksByUser(db *gorm.DB, userID uuid.UUID) ([]models.Task, error) {
	var tasks []models.Task

	if err := db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// UpdateTask implements [TaskService].
func (t *TaskServiceImpl) UpdateTask(db *gorm.DB, taskID uuid.UUID,
	userID uuid.UUID, update models.TaskUpdate) (
	models.Task, error) {
	updates := map[string]interface{}{}
	if update.Title != nil {
		updates["title"] = *update.Title
	}
	if update.Description != nil {
		updates["description"] = *update.Description
	}
	if update.Status != nil {
		updates["status"] = *update.Status
	}

	// Nothing to change. still verify the task exists
	if len(updates) == 0 {
		var task models.Task
		if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
			return models.Task{}, err
		}
		return task, nil
	}

	result := db.Model(&models.Task{}).Where("id = ? AND user_id = ?", taskID, userID).Updates(updates)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	if result.RowsAffected == 0 {
		return models.Task{}, gorm.ErrRecordNotFound
	}

	var task models.Task
	if err := db.Where("id = ?", taskID).First(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}
