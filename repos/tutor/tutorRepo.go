package tutor

import (
	"course_management/db"
	"course_management/models"
)

type TutorRepo struct{}

func (t TutorRepo) Create(tutor *models.Tutor) error {
	result := db.DB.Create(&tutor)
	return result.Error
}

func (t TutorRepo) GetTutorByEmail(email string) (models.Tutor, error) {
	tutor := models.Tutor{}
	result := db.DB.Find(&tutor, "email = ?", email)
	return tutor, result.Error
}
