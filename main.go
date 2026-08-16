package main

import (
	"log"
	"task-mgmt/database"
	"task-mgmt/middleware"
	"task-mgmt/routes"

	"github.com/gin-gonic/gin"
)

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
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.RateLimiter())

	routes.SetupRoutes(r, db)

	r.Run(":8080")
}
