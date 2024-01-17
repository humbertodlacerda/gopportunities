package movement

import (
	"fmt"
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdateMovement(context echo.Context) error {
	request := UpdateMovementRequest{}

	context.Bind(&request)

	err := request.Validate()

	if err != nil {
		handler.Logger.Errorf("validation error: %v", err.Error())
		return handler.SendError(context, http.StatusBadRequest, err.Error())
	}

	id := context.QueryParam("id")

	if id == "" {
		return handler.SendError(context, http.StatusBadRequest, handler.ErrorParamIsRequired(id, "queryParameter").Error())
	}

	movement := schemas.Movement{}

	err = handler.Db.First(&movement, id).Error

	if err != nil {
		return handler.SendError(context, http.StatusNotFound, fmt.Sprintf("movement with id: %s not found", id))
	}

	if request.Date != "" {
		movement.Date = request.Date
	}

	if request.Description != "" {
		movement.Description = request.Description
	}

	if request.Category != "" {
		movement.Category = request.Category
	}

	if request.Account != "" {
		movement.Account = request.Account
	}

	if request.Value > 0 {
		movement.Value = request.Value
	}

	if request.Status > 0 {
		movement.Status = request.Status
	}

	if err := handler.Db.Save(&movement).Error; err != nil {
		handler.Logger.Errorf("error updating movement: %v", err.Error())
		return handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error updating movement with id: %s", id))
	}

	return handler.SendSuccess(context, "update-movement", movement)
}
