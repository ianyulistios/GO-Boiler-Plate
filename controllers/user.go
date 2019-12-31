package controllers

import (
	"GO-Boiler-Plate/config"
	"GO-Boiler-Plate/helpers"
	"GO-Boiler-Plate/services"

	"github.com/fatih/structs"

	"github.com/labstack/echo"
)

//UserController Represented of User Service Instance
type UserController struct {
	userSV services.UserService
}

func (userSV *UserController) getUser(c echo.Context) (err error) {
	data, err := userSV.userSV.Get()
	if err != nil {
		errResponse, errResponseCode := helpers.ErrResponseHelper(err, "ERR_GENERAL")
		return c.JSONPretty(errResponseCode, errResponse, config.JSONString)
	}
	mapInterfaceData := structs.Map(data)
	successResponse, successresponsecode := helpers.SuccessResponseHelper(mapInterfaceData, "SUCCESS_200")
	return c.JSONPretty(successresponsecode, successResponse, config.JSONString)
}
