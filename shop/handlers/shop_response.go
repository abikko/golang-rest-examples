package handlers

import "github.com/labstack/echo/v4"

type baseResponse struct {
	Message string `json:"message"`
}

func response(context echo.Context, responseCode int, message string) error {
	return context.JSON(responseCode, &baseResponse{Message: message})
}
