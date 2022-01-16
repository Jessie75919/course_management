package validators

import (
	"course_management/repos/tutor"
	"github.com/go-playground/validator/v10"
)

var Existed validator.Func = func(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	tutor, _ := new(tutor.TutorRepo).GetTutorByEmail(email)
	if tutor.ID == 0 {
		return true
	}
	return false
}
