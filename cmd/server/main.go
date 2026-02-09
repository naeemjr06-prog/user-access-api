package main

import (
	"github.com/naeemjr06-prog/user-access-api/pkg/config"
	"github.com/naeemjr06-prog/user-access-api/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	database.Connect()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/health/db", func(c *gin.Context) {
		sqlDB, _ := database.DB.DB()
		if err := sqlDB.Ping(); err != nil {
			c.JSON(500, gin.H{"db": "down"})
			return
		}
		c.JSON(200, gin.H{"db": "up"})
	})

	r.Run(":8080")
}
