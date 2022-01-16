package routes

import (
	controllers "course_management/controllers/tutors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		tutors := v1.Group("/tutors")
		{
			tutorController := new(controllers.TutorController)
			tutors.POST("/sign-up", tutorController.SignUp)
			tutors.POST("/login", tutorController.Login)
			tutors.GET("/profile", tutorController.Profile)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})
}
