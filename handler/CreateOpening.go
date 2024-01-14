package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"net/http"
)

func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Link:     request.Link,
		Location: request.Location,
		Remote:   *request.Remote,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(context, "create-opeing", opening)
	return
}
