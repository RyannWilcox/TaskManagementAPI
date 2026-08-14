package handlers

import (
	"net/http"

	"task-mgmt/models"
	"task-mgmt/services"
	"task-mgmt/utils"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type TaskHandler struct {
	db          *gorm.DB
	taskService services.TaskService
}

func NewTaskHandler(db *gorm.DB, taskService services.TaskService) *TaskHandler {
	return &TaskHandler{db: db, taskService: taskService}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.Error(err)
		return
	}

	if err := h.taskService.CreateTask(h.db, task); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, utils.MessageResponse{
		Message: "task created successfully",
	})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	// Implement task update logic here

	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {

	c.JSON(http.StatusNoContent, nil)
}

// Retrieve a task by a provided ID
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	taskID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.Error(err)
	}

	// Get the currently logged in user off the context.
	userID := c.MustGet("userID").(uuid.UUID)

	task, err := h.taskService.GetTaskByID(h.db, taskID, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) GetTasksByUser(c *gin.Context) {

	c.JSON(http.StatusOK, "tasks list by user")
}

func (h *TaskHandler) GetTasks(c *gin.Context) {

	c.JSON(http.StatusOK, "tasks list")
}
