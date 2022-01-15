package db

import (
	"course_management/config"
	"course_management/models"
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

	DB = db

	runMigrate(db)

	// Migrate the schema
	//db.AutoMigrate(&Product{})

}

func runMigrate(db *gorm.DB) {
}
