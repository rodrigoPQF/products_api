package handlers

import (
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewSuccessResponse(ctx echo.Context, statusCode int, message string, payload interface{}) error {
	return ctx.JSON(statusCode, BaseResponse{
		Status:  true,
		Message: message,
		Payload: payload,
	})
}

func NewErrorReponse(ctx echo.Context, statusCode int, err string) error {
	return ctx.JSON(statusCode, BaseResponse{
		Status:  false,
		Message: err,
	})
}
