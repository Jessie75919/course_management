package validators

import (
	"course_management/repos/tutor"
	"github.com/go-playground/validator/v10"
)

var EmailExist validator.Func = func(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	notExisted := fl.Param() == "false"

	tutor, _ := new(tutor.TutorRepo).GetTutorByEmail(email)
	if tutor.ID == 0 {
		return notExisted
	}
	return !notExisted
}
