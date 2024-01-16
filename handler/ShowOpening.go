package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"net/http"
)

func ShowOpeningHandler(context *gin.Context) {
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

	sendSuccess(context, "show-opening", opening)
}
