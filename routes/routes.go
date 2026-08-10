package routes

import (
	"task-mgmt/handlers"
	"task-mgmt/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	authHandler := handlers.NewAuthHandler(db, services.NewAuthService())
	taskHandler := handlers.NewTaskHandler(db, nil)
	refreshHandler := handlers.NewRefreshHandler(db, services.NewAuthService())
	registerHandler := handlers.NewRegisterHandler(db, services.NewRegisterService())
	userHandler := handlers.NewUserHandler(db, nil)

	v1 := router.Group("/api/v1")
	{
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/register", registerHandler.Registration)
			authRoutes.POST("/login", authHandler.Token)
			authRoutes.POST("/refresh", refreshHandler.Refresh)
		}
		taskRoutes := v1.Group("/tasks")
		{
			taskRoutes.POST("", taskHandler.CreateTask)
			taskRoutes.PUT("/:id", taskHandler.UpdateTask)
			taskRoutes.DELETE("/:id", taskHandler.DeleteTask)
			taskRoutes.GET("/:id", taskHandler.GetTaskByID)
			taskRoutes.GET("", taskHandler.GetTasks)
		}
		userRoutes := v1.Group("/users")
		{
			userRoutes.DELETE("/:user_id", userHandler.DeleteUser)
			userRoutes.GET("", userHandler.GetUsers)
			userRoutes.GET("/:user_id/tasks", taskHandler.GetTasksByUser)
			userRoutes.GET("/profile", userHandler.GetUserProfile)
			userRoutes.GET("/profile/:user_id", userHandler.GetUserProfileByUserId)
		}
	}

}
