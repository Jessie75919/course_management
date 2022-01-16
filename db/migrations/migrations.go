package migrations

import (
	"course_management/models"
	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Tutor{})
	db.AutoMigrate(&models.Student{})
	db.AutoMigrate(&models.Lesson{})
	db.AutoMigrate(&models.Transaction{})
	db.AutoMigrate(&models.Fee{})
	db.AutoMigrate(&models.Wallet{})
	db.AutoMigrate(&models.Topic{})
}
