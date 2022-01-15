package models

import (
	"gorm.io/gorm"
	"time"
)

type Tutor struct {
	ID          uint
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

type Lesson struct {
	ID           uint
	TutorID      uint
	StudentID    uint
	PeriodAmount uint
	LessonTime   time.Time
	Note         string `gorm:"type:text"`
	Transaction  Transaction
	gorm.Model
}

type Transaction struct {
	ID                   uint
	TutorID              uint
	StudentID            uint
	LessonID             uint
	OriginalWalletAmount uint
	Amount               uint
	FinalWalletAmount    uint
	TransactionDate      time.Time
	gorm.Model
}

type Wallet struct {
	ID        uint
	TutorID   uint
	StudentID uint
	gorm.Model
}

type Fee struct {
	ID        uint
	TutorID   uint
	StudentID uint
	BeginDate time.Time
	EndDate   time.Time
	Amount    uint
	gorm.Model
}

type Topic struct {
	ID      uint
	Name    string
	Student []Student `gorm:"many2many:student_topic;"`
	gorm.Model
}
