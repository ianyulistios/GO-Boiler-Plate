package helpers

import (
	"GO-Boiler-Plate/config"
)

var databaseConfig = config.JSONConstants("db")

//Database represented to get configuration in constants folder of db file
func Database(key string) interface{} {
	var result = databaseConfig.Get(key)
	return result
}
