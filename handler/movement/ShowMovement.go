package movement

import (
	"fmt"
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ShowMovement(context echo.Context) error {
	id := context.QueryParam("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrorParamIsRequired(id, "queryParameter").Error())
		return nil
	}

	movement := schemas.Movement{}

	if err := handler.Db.First(&movement, id).Error; err != nil {
		return handler.SendError(context, http.StatusNotFound, fmt.Sprintf("movement with id: %s not found", id))
	}

	return handler.SendSuccess(context, "show-movement", movement)
}
