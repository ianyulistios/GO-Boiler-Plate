package controllers

import (
	"GO-Boiler-Plate/services"

	"github.com/labstack/echo"
)

//UserService Represented of User Service Instance
type UserService struct {
	service services.UserService
}

func (service *UserService) getUser(c echo.Context) {

}
