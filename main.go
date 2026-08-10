package main

import (
	"log"
	"task-mgmt/database"
	"task-mgmt/middleware"
	"task-mgmt/routes"
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://host.docker.internal"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes(r, db)

	r.Run(":8080")
}
