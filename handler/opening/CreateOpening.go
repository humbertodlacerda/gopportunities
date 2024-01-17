package opening

import (
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateOpening(context echo.Context) error {
	request := CreateOpeningRequest{}

	context.Bind(&request)

	err := request.Validate()

	if err != nil {
		handler.Logger.Errorf("error creating opening: %v", err.Error())
		return handler.SendError(context, http.StatusBadRequest, err.Error())
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Link:     request.Link,
		Location: request.Location,
		Remote:   *request.Remote,
		Salary:   request.Salary,
	}

	err = handler.Db.Create(&opening).Error

	if err != nil {
		handler.Logger.Errorf("error creating opening: %v", err.Error())
		return handler.SendError(context, http.StatusInternalServerError, "error creating opening on database")
	}

	return handler.SendSuccess(context, "create-opening", opening)
}
