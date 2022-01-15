package main

import (
	"course_management/config"
	"course_management/db"
	"course_management/logger"
	"github.com/gin-gonic/gin"
)

func setup() {
	config.Setup()
	logger.Setup()
	db.Setup()

	logger.Debug.Println("test something log")
}

func main() {
	setup()
}

