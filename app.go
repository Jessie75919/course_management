package main

import (
	"course_management/config"
	"course_management/db"
	"course_management/gin_setting"
	"course_management/logger"
)

func setup() {
	config.Setup()
	logger.Setup()
	db.Setup()

	gin_setting.GinSetup()
	logger.Debug.Println("setup was done !")
}

func main() {
	setup()
}
