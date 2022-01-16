package models

import (
	"gorm.io/gorm"
	"time"
)

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
