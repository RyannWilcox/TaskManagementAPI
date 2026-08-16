package main

import (
	"log"
	"task-mgmt/database"
	"task-mgmt/middleware"
	"task-mgmt/routes"

	"github.com/gin-gonic/gin"
)

// @title Task Management API
// @version 1.0
// @description This is a backend API for a task management system
// @host localhost:8080
// @BasePath /api/v1

func main() {
	dbConfig := database.NewDatabaseConfig()
	db, err := dbConfig.Connect()

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	defer sqlDB.Close()

	r := gin.Default()
	r.Use(middleware.SecurityHeaders())

	routes.SetupRoutes(r, db)

	r.Run(":8080")
}
