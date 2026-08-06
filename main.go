package main

import (
	"log"
	"task-mgmt/database"

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

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	r.Run(":8080")
}
