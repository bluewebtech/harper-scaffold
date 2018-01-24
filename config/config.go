package config

import (
	"github.com/joho/godotenv"
	"log"
)

var Configuration = map[string]interface{}{}

var setEnvironmentVariables = BootEnvironmentVariables()

func BootEnvironmentVariables() bool {
	load := godotenv.Load()
	var status = true

	if load != nil {
		log.Fatal("Error loading .env file")
		status = false
	}

	return status
}
