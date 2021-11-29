package main

import (
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo/v4"
)

func Service1() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
	e.POST("/battle", printBattle)
	e.GET("/users/:id", getUser)
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func printBattle(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
