package main

import (
	"net/http"

	"github.com/avenue-golang/statement/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database connection
	dbConn := db.Init()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		if dbConn.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "database": "disconnected"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "database": "connected"})
	})

	// Other API endpoints (placeholder)
	r.GET("/events", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Events endpoint"})
	})

	r.Run(":8080")
}
