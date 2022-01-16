package models

import "gorm.io/gorm"

type Student struct {
	ID          uint
	Name        string `gorm:"not null"`
	Username    string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Password    string `gorm:"not null"`
	TutorID     uint
	Lesson      []Lesson
	Transaction []Transaction
	Wallet      Wallet
	Fee         []Fee
	Topic       []Topic `gorm:"many2many:student_topic;"`
	gorm.Model
}
