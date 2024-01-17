package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SendError(context echo.Context, code int, msg string) error {
	context.Response().Header().Set("Content-type", "application/json")

	return context.JSON(code, echo.Map{
		"message":   msg,
		"errorCode": code,
	})
}

func SendSuccess(context echo.Context, operation string, data interface{}) error {
	context.Response().Header().Set("Content-type", "application/json")

	return context.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("operation from handler: %s successfull", operation),
		"data":    data,
	})
}
