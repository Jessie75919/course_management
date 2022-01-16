package tutor

type SignupTutorRequest struct {
	Name              string `json:"name" binding:"required"`
	Username          string `json:"username" binding:"required"`
	Email             string `json:"email" binding:"required,email,existed"`
	Password          string `json:"password" binding:"required"`
	ConfirmedPassword string `json:"confirmed_password" binding:"required,eqfield=Password"`
}
