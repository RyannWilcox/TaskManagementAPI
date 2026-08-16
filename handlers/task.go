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

// CreateTask godoc
// @Summary      Create a task
// @Description  Creates a new task
// @Tags         tasks
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      models.Task  true  "Task to create"
// @Success      201      {object}  models.Task
// @Failure      400      {object}  utils.HTTPError  "invalid request payload"
// @Router       /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.Error(err)
		return
	}

	task.UserID = c.MustGet("userID").(uuid.UUID)

	newTask, err := h.taskService.CreateTask(h.db, task)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// UpdateTask godoc
// @Summary      Partially update a task
// @Description  Updates one or more of title, description, and status on a task owned by the authenticated user.
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      string            true  "Task ID"
// @Param        request  body      models.TaskUpdate  true  "Fields to update (all optional)"
// @Success      200      {object}  models.Task
// @Failure      400      {object}  utils.HTTPError  "invalid task id or invalid request payload"
// @Failure      404      {object}  utils.HTTPError  "record not found"
// @Router       /tasks/{id} [patch]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.Error(utils.ErrInvalidUUID)
		return
	}

	var update models.TaskUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.Error(err)
		return
	}

	// Get the currently logged in user off the context.
	userID := c.MustGet("userID").(uuid.UUID)

	updatedTask, err := h.taskService.UpdateTask(h.db, taskID, userID, update)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask godoc
// @Summary      Delete a task
// @Description  Deletes a task by ID
// @Tags         tasks
// @Security     BearerAuth
// @Produce      json
// @Param        id  path      string  true  "Task ID"
// @Success      204 {object}  utils.MessageResponse
// @Failure      400 {object}  utils.HTTPError  "invalid task id"
// @Failure      404 {object}  utils.HTTPError  "record not found"
// @Router       /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.Error(utils.ErrInvalidUUID)
		return
	}

	if err := h.taskService.DeleteTask(h.db, taskID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetTaskByID godoc
// @Summary      Get a task by ID
// @Description  Returns a single task owned by the authenticated user
// @Tags         tasks
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "Task ID"
// @Success      200  {object}  models.Task
// @Failure      400  {object}  utils.HTTPError  "invalid task id"
// @Failure      404  {object}  utils.HTTPError  "record not found"
// @Router       /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	taskID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.Error(utils.ErrInvalidUUID)
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

// GetTasksByUser godoc
// @Summary      List a user's tasks
// @Description  Returns all tasks belonging to the authenticated user
// @Tags         tasks
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}  models.Task
// @Router       /users/{user_id}/tasks [get]
func (h *TaskHandler) GetTasksByUser(c *gin.Context) {
	// Get the currently logged in user off the context.
	userID := c.MustGet("userID").(uuid.UUID)

	tasks, err := h.taskService.GetTasksByUser(h.db, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTasks godoc
// @Summary      List tasks
// @Description  Returns tasks
// @Tags         tasks
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}  models.Task
// @Router       /tasks [get]
func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.taskService.GetTasks(h.db)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}
