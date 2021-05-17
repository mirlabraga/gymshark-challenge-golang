package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mirlabraga/gymshark-challenge-golang/src/services"
)

func getOrders(c echo.Context) error {
	packages := services.Calculation(250)
	return c.JSON(http.StatusCreated, packages)
}

func main() {

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.POST("/orders", getOrders)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}
