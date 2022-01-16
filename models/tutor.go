package models

import (
	"gorm.io/gorm"
)

type Tutor struct {
	ID          int
	Name        string `gorm:"not null"`
	Username    string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Students    []Student
	Lesson      []Lesson
	Transaction []Transaction
	Wallet      Wallet
	Fee         []Fee
	gorm.Model
}
