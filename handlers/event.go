package handlers

import (
	"github.com/gin-gonic/gin"
	"project.com/event-booking/service"
)

func Apis(server *gin.Engine) {

	// If a function is resgitered as a handler, gin "AUTOMATICALLY"
	// sends the gin.Context parameter to the function
	// In this case, getEvent() function is used as a handler in the server.GET() request,
	// therefore it automatically sends the pointer to the gin.Context to the getEvents function, and we must
	// recieve it in the definition of our function
	// context will be set by the gin automatically

	server.GET("/events", service.GetEvents)
	server.GET("/event/:id", service.GetEvent)
	server.POST("/create/event", service.CreateEvent)
	server.Run(":8080")
}
