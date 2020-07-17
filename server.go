package main

import (
	handlers "crud-go/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/get", handlers.GetPing)

	e.Logger.Fatal(e.Start(":8080"))
}
