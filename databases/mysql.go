package databases

import (
	"GO-Boiler-Plate/helpers"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gormInstance *gorm.DB //instance of DB of mysql

//InitDBMysql Represented of initalize db connection to Mysql
func InitDBMysql() (*gorm.DB, error) {
	APPENV := os.Getenv("APP_ENV")
	configConst := helpers.Database(APPENV + "." + "Mysql")
	mapConfigConst := configConst.(map[string]interface{})
	fmt.Print("disini")
	fmt.Print(mapConfigConst)
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mapConfigConst["username"].(string),
		mapConfigConst["password"].(string),
		mapConfigConst["host"].(string),
		mapConfigConst["port"].(string),
		mapConfigConst["db"].(string),
	)
	fmt.Print(connString)
	return gorm.Open("mysql", connString)
}

//GetConnection represented to get the connction based on singleton pattern
func GetConnection() *gorm.DB {
	if gormInstance == nil {
		conn, err := InitDBMysql()
		if err != nil {
			panic(fmt.Errorf("Error Mysql Connection: %s", err))
		}
		gormInstance = conn
	}
	return gormInstance
}
