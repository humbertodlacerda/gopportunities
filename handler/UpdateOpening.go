package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"net/http"
)

func UpdateOpeningHandler(context *gin.Context) {
	request := UpdateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")

	if id == "" {
		sendError(context, http.StatusBadRequest, errorParamIsRequired(id, "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
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

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("error updating opening with id: %s", id))
		return
	}

	sendSuccess(context, "update-opening", opening)
}
