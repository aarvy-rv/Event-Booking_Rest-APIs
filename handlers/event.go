package handlers

import (
	"github.com/gin-gonic/gin"
	"project.com/event-booking/middlewares"
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

	// We can group end points so that it can use mididdleware defined once
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/create/event", service.CreateEvent)
	authenticated.PUT("/event/update/:id", service.UpdateEvent)
	authenticated.DELETE("/event/delete/:id", service.Delete)
	authenticated.POST("/events/:id/register", service.RegisterForEvent)
	authenticated.DELETE("/events/:id/register/cancel", service.CancelRegistration)

	//server.POST("/create/event", middlewares.Authenticate,service.CreateEvent)
	//server.PUT("/event/update/:id", service.UpdateEvent)
	//server.DELETE("/event/delete/:id", service.Delete)

	server.POST("/signup", service.Signup)
	server.POST("/login", service.Login)
	server.Run(":8080")
}
