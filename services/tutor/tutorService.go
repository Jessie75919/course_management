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
		Email:    r.Email,
		Password: bcrypt.HashPassword(r.Password),
	}

	err := new(repo.TutorRepo).Create(&tutor)
	if err != nil {
		return -1, err
	}

	return tutor.ID, nil
}

func Login(r *request.LoginTutorRequest) error {
	tutor := models.Tutor{}
	tutor, err := new(repo.TutorRepo).GetTutorByEmail(r.Email)

	if err != nil {
		return err
	}

	return bcrypt.ComparePW(r.Password, tutor.Password)
}
