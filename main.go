package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
          return c.JSON(http.StatusOK, &echo.Map{"data": "Hello from Echo and mongoDB"})
    })

    e.Logger.Fatal(e.Start(":6060"))
}