package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//JSONString Represented of Global Variabel to define JSON Pretty string argument
var JSONString = ""

//JSONConstants Represented of function to get Json Data in Constants Folder
func JSONConstants(constantName string) *viper.Viper {
	jsonConstant := viper.New()
	jsonConstant.SetConfigName(constantName)
	jsonConstant.AddConfigPath("constants")
	errJSONConstant := jsonConstant.ReadInConfig()
	if errJSONConstant != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", errJSONConstant))
	}
	return jsonConstant
}
