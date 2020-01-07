package controllers

import (
	"GO-Boiler-Plate/config"
	"GO-Boiler-Plate/helpers"
	"GO-Boiler-Plate/services"

	"github.com/fatih/structs"

	"github.com/labstack/echo"
)

//UserSV Represented of User Service Instance
var userSV *services.UserService

// @title GetUser
// @description get User Data
// @success 200 {string} string	"success"
// @failure 500 {string} string	"error"
// @router /user [get]
//GetUser Represented to Get Data User From DB
func GetUser(c echo.Context) (err error) {
	data, err := userSV.Get()
	if err != nil {
		errResponse, errResponseCode := helpers.ErrResponseHelper(err, "ERR_GENERAL")
		return c.JSONPretty(errResponseCode, errResponse, config.JSONString)
	}
	mapInterfaceData := structs.Map(data)
	successResponse, successresponsecode := helpers.SuccessResponseHelper(mapInterfaceData, "SUCCESS_200")
	return c.JSONPretty(successresponsecode, successResponse, config.JSONString)
}
