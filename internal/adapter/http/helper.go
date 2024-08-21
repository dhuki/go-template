package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(c echo.Context, status int, err error) error {
	res := Response{
		Status:  "ERROR",
		Message: err.Error()}
	return c.JSON(status, res)
}

func ResponseSuccess(c echo.Context, msg string, r interface{}) error {
	res := Response{
		Status:  "OK",
		Message: msg,
		Data:    r,
	}
	return c.JSON(http.StatusOK, res)
}
