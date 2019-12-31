package routes

import (
	"GO-Boiler-Plate/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// API Represented of function to Define Routes
func API(e *echo.Echo) {
	route := e.Group("/v1")
	route.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	route.GET("/user", controllers.GetUser)
}
