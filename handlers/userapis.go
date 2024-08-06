package handlers

import (
	"github.com/gin-gonic/gin"
	"project.com/event-booking/service"
)

func Uapis(server *gin.Engine) {

	server.POST("/signup", service.Signup)
}
