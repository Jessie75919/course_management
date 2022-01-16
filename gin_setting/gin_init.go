package gin_setting

import (
	"course_management/gin_setting/routes"
	"course_management/gin_setting/validators"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
	"os"
)

func GinSetup() {
	r := gin.Default()

	setupValidators()

	setupLog()

	routes.SetupRoutes(r)

	r.Run(":80")
}

func setupValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("existed", validators.Existed)
	}
}

func setupLog() {
	f, _ := os.Create("var/log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}
