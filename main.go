package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sautsihotang/DatingApp/config"
	"github.com/sautsihotang/DatingApp/routes"
)

func main() {
	config.Init()

	e := echo.New()

	routes.Route(e)

	e.Logger.Fatal(e.Start(":8080"))
}
