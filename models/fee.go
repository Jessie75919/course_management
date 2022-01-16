package models

import (
	"gorm.io/gorm"
	"time"
)

type Fee struct {
	ID        uint
	TutorID   uint
	StudentID uint
	BeginDate time.Time
	EndDate   time.Time
	Amount    uint
	gorm.Model
}
