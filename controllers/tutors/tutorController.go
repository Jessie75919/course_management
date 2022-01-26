package controllers

import (
	request "course_management/requests/tutor"
	service "course_management/services/tutor"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TutorController struct{}

func (tc *TutorController) SignUp(c *gin.Context) {
	var data request.SignupTutorRequest
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	id, err := service.SignUp(&data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "sign up failed, try again"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sign up successfully. id: " + strconv.Itoa(id)})
}

func (tc *TutorController) Login(c *gin.Context) {
	var data request.LoginTutorRequest
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	token, err := service.Login(&data)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login in failed, try again"})
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"access_token": token,
		},
	)
}
}
