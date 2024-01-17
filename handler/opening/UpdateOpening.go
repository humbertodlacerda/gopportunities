package opening

import (
	"fmt"
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdateOpening(context echo.Context) error {
	request := UpdateOpeningRequest{}

	context.Bind(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("validation error: %v", err.Error())
		return handler.SendError(context, http.StatusBadRequest, err.Error())
	}

	id := context.QueryParam("id")

	if id == "" {
		return handler.SendError(context, http.StatusBadRequest, handler.ErrorParamIsRequired(id, "queryParameter").Error())
	}

	opening := schemas.Opening{}

	if err := handler.Db.First(&opening, id).Error; err != nil {
		return handler.SendError(context, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := handler.Db.Save(&opening).Error; err != nil {
		handler.Logger.Errorf("error updating opening: %v", err.Error())
		return handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error updating opening with id: %s", id))
	}

	return handler.SendSuccess(context, "update-opening", opening)
}
