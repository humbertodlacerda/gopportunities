package movement

import (
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ListMovement(context echo.Context) error {
	movements := []schemas.Movement{}

	err := handler.Db.Find(&movements).Error

	if err != nil {
		return handler.SendError(context, http.StatusInternalServerError, "error listing movements")
	}

	return handler.SendSuccess(context, "list-movements", movements)
}
