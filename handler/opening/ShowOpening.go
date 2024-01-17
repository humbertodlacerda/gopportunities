package opening

import (
	"fmt"
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ShowOpening(context echo.Context) error {
	id := context.QueryParam("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrorParamIsRequired(id, "queryParameter").Error())
		return nil
	}

	opening := schemas.Opening{}

	if err := handler.Db.First(&opening, id).Error; err != nil {
		return handler.SendError(context, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
	}

	return handler.SendSuccess(context, "show-opening", opening)
}
