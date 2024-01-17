package router

import (
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/handler/movement"
	"github.com/humbertodlacerda/gopportunities/handler/opening"
	"github.com/labstack/echo/v4"
)

func initializeRoutes(router *echo.Echo) {
	handler.InitializeHandler()
	openingGroup := router.Group("/api")
	{
		openingGroup.GET("/opening", opening.ShowOpening)
		openingGroup.POST("/opening", opening.CreateOpening)
		openingGroup.DELETE("/opening", opening.DeleteOpening)
		openingGroup.PUT("/opening", opening.UpdateOpening)
		openingGroup.GET("/openings", opening.ListOpenings)
	}

	movementGroup := router.Group("/api")
	{
		movementGroup.GET("/movement", movement.ShowMovement)
		movementGroup.POST("/movement", movement.CreateMovement)
		movementGroup.DELETE("/movement", movement.DeleteMovement)
		movementGroup.PUT("/movement", movement.UpdateMovement)
		movementGroup.GET("/movements", movement.ListMovement)
	}
}
