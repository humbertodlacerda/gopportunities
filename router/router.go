package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Initialize() {
	// Initialize router
	//router := gin.Default()
	router := echo.New()
	router.Use(corsMiddleware())

	// Initialize routes
	initializeRoutes(router)

	// Run the server
	err := router.Start(":8080")
	if err != nil {
		return
	}
}

func corsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if c.Request().Method == http.MethodOptions {
				return c.NoContent(http.StatusOK)
			}
			// Continuar para o próximo middleware ou rota
			return next(c)
		}
	}
}
