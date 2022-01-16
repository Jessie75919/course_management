package models

import "gorm.io/gorm"

type Topic struct {
	ID      uint
	Name    string
	Student []Student `gorm:"many2many:student_topic;"`
	gorm.Model
}
