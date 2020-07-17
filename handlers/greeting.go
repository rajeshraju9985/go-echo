package handlers

import (
	"net/http"
	"time"

	model "crud-go/model"

	"github.com/labstack/echo/v4"
)

func GetPing(c echo.Context) (err error) {

	greeting := &model.InlineResponse200{}
	greeting.Greeting = "Poker Service Up"
	greeting.Date = time.Now().String()

	return c.JSON(http.StatusOK, greeting)
}
