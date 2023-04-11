package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Success struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func SuccessResponse(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, &Success{
		Status: http.StatusOK,
		Data:   data,
	})
}

func BadResponse(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusOK, &Error{
		Status:  http.StatusBadRequest,
		Message: message,
	})
}
