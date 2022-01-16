package models

import "gorm.io/gorm"

type Wallet struct {
	ID        uint
	TutorID   uint
	StudentID uint
	gorm.Model
}
