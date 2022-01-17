package tutor

type SignupTutorRequest struct {
	Name              string `json:"name" binding:"required"`
	Email             string `json:"email" binding:"required,email,emailExist=false"`
	Password          string `json:"password" binding:"required"`
	ConfirmedPassword string `json:"confirmed_password" binding:"required,eqfield=Password"`
}
