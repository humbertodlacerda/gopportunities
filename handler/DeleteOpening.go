package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"net/http"
)

func DeleteOpeningHandler(context *gin.Context) {
	id := context.Query("id")

	if id == "" {
		sendError(context, http.StatusBadRequest, errorParamIsRequired(
			"id",
			"queryParameter",
		).Error())
		return
	}

	opening := schemas.Opening{}

	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	// Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf(
			"error deleeting opening with id: %s",
			id,
		))
		return
	}

	sendSuccess(context, "delete-opening", opening)
}
