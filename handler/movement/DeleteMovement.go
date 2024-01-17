package movement

import (
	"fmt"
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DeleteMovement(context echo.Context) error {
	id := context.QueryParam("id")

	if id == "" {
		return handler.SendError(context, http.StatusBadRequest, handler.ErrorParamIsRequired(
			"id",
			"queryParameter",
		).Error())
	}

	movement := schemas.Movement{}

	err := handler.Db.First(&movement, id).Error

	if err != nil {
		return handler.SendError(context, http.StatusNotFound, fmt.Sprintf("movement with id: %s not found", id))
	}

	err = handler.Db.Delete(&movement).Error

	if err != nil {
		return handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf(
			"error deleeting movement with id: %s",
			id,
		))
	}

	return handler.SendSuccess(context, "delete-movement", movement)
}
