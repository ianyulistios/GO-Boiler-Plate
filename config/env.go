package config

import "os"

//Env Represented of function to set environment
func Env() {
	os.Setenv("APP_ENV", "DEVELOPMENT")
}
