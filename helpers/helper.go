package helpers

import (
	"GO-Boiler-Plate/config"
)

var databaseConfig = config.JSONConstants("db")
var responseConfig = config.JSONConstants("response")

//Database represented to get configuration in constants folder of db file
func Database(key string) interface{} {
	result := databaseConfig.Get(key)
	return result
}

// GetJSONResponse Represnted of Function to Get The Json Structure of Response
func GetJSONResponse(key string) interface{} {
	result := responseConfig.Get(key)
	return result
}
