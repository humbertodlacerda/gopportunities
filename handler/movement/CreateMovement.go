package movement

import (
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateMovement(context echo.Context) error {
	request := CreateMovementRequest{}

	context.Bind(&request)

	err := request.Validate()

	if err != nil {
		handler.Logger.Errorf("error creating movement: %v", err.Error())
		return handler.SendError(context, http.StatusBadRequest, err.Error())
	}

	movement := schemas.Movement{
		Date:        request.Date,
		Description: request.Description,
		Category:    request.Category,
		Account:     request.Account,
		Value:       request.Value,
		Status:      request.Status,
	}

	err = handler.Db.Create(&movement).Error
	if err != nil {
		handler.Logger.Errorf("error creating movement: %v", err.Error())
		return handler.SendError(context, http.StatusInternalServerError, "error creating on database")
	}

	return handler.SendSuccess(context, "create-movement", movement)
}
