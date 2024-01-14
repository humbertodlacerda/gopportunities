package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Initialize() {
	// Initialize router
	router := gin.Default()

	router.Use(corsMiddleware())

	// Initialize routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Continuar para o próximo middleware ou rota
		c.Next()
	}
}
