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
	taskID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	var update models.TaskUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.Error(err)
		return
	}

	userID := c.MustGet("userID").(uuid.UUID)

	updatedTask, err := h.taskService.UpdateTask(h.db, taskID, userID, update)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.taskService.DeleteTask(h.db, taskID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, utils.MessageResponse{
		Message: "task deleted",
	})
}

// Retrieve a task by a provided ID
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	taskID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
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
	userID := c.MustGet("userID").(uuid.UUID)

	tasks, err := h.taskService.GetTasksByUser(h.db, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.taskService.GetTasks(h.db)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, tasks)
}
