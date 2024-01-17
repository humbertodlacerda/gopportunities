package router

import (
	"github.com/humbertodlacerda/gopportunities/handler"
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

	//movementGroup := router.Group("/api")
	//{
	//	movementGroup.GET("/movement")
	//	movementGroup.POST("/movement")
	//	movementGroup.DELETE("/movement")
	//	movementGroup.PUT("/movement")
	//	movementGroup.GET("/movements")
	//}
}
