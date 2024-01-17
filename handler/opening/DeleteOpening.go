package opening

import (
	"fmt"
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DeleteOpening(context echo.Context) error {
	id := context.QueryParam("id")

	if id == "" {
		return handler.SendError(context, http.StatusBadRequest, handler.ErrorParamIsRequired(
			"id",
			"queryParameter",
		).Error())
	}

	opening := schemas.Opening{}

	// Find Opening
	if err := handler.Db.First(&opening, id).Error; err != nil {
		return handler.SendError(context, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
	}
	// Delete Opening
	if err := handler.Db.Delete(&opening).Error; err != nil {
		return handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf(
			"error deleeting opening with id: %s",
			id,
		))
	}

	return handler.SendSuccess(context, "delete-opening", opening)
}
