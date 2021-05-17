package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mirlabraga/gymshark-challenge-golang/src/models"
	"github.com/mirlabraga/gymshark-challenge-golang/src/services"
)

func getOrders(c echo.Context) error {

	order := &models.Order{}

	if err := c.Bind(order); err != nil {
		return err
	}

	packages := services.Calculation(order.Quantity)
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
