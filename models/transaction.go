package models

import (
	"gorm.io/gorm"
	"time"
)

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
