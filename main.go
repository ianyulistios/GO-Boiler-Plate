package main

import (
	"GO-Boiler-Plate/config"
	"GO-Boiler-Plate/routes"

	"github.com/labstack/echo"
)

func main() {
	config.Env()
	e := echo.New()
	routes.API(e)
	e.Logger.Fatal(e.Start(":1412"))
}
