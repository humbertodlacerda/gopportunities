package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"net/http"
)

func ListOpeningsHandler(context *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(context, "list-opening", openings)
}
