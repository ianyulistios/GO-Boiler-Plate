package main

import (
	"GO-Boiler-Plate/config"
	"GO-Boiler-Plate/helpers"
	"GO-Boiler-Plate/migration/tables"
	"GO-Boiler-Plate/routes"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	config.Env()
	e := echo.New()
	switch helpers.Command() {
	case "migrate":
		tables.ExecuteMigration()
		break
	default:
		fmt.Println("Invalid command")
	}
	routes.API(e)
	if helpers.Command() != "migrate" {
		e.Logger.Fatal(e.Start(":1412"))
	}
}
