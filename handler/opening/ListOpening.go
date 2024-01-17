package opening

import (
	"github.com/humbertodlacerda/gopportunities/handler"
	"github.com/humbertodlacerda/gopportunities/schemas"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ListOpenings(context echo.Context) error {
	openings := []schemas.Opening{}

	if err := handler.Db.Find(&openings).Error; err != nil {
		return handler.SendError(context, http.StatusInternalServerError, "error listing openings")
	}

	return handler.SendSuccess(context, "list-opening", openings)
}
