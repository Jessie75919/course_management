package tutor

import (
	"course_management/helpers/bcrypt"
	"course_management/models"
	repo "course_management/repos/tutor"
	request "course_management/requests/tutor"
)

func SignUp(r *request.SignupTutorRequest) (int, error) {
	tutor := models.Tutor{
		Name:     r.Name,
		Username: r.Username,
		Email:    r.Email,
		Password: bcrypt.HashPassword(r.Password),
	}

	err := new(repo.TutorRepo).Create(&tutor)
	if err != nil {
		return -1, err
	}

	return tutor.ID, nil
}
