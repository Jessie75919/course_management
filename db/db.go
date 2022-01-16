package db

import (
	"course_management/config"
	"course_management/db/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Setup() {
	dsn := config.EnvConfig.DBUrl
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	migrations.RunMigrate(db)

	DB = db
}
