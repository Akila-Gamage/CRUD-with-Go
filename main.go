package main

import (
	"crud-with-go/configs"
	"crud-with-go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    //run database
    configs.ConnectDB()

    //routes
    routes.UserRoute(e)

    e.Logger.Fatal(e.Start(":6050"))
}