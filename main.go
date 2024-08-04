package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)
	// If a function is resgitered as a handler, gin "AUTOMATICALLY"
	// sends the gin.Context parameter to the function
	// In this case, getEvent() function is used as a handler in the server.GET() request,
	// therefore it automatically sends the pointer to the gin.Context to the getEvents function, and we must
	// recieve it in the definition of our function
	// context will be set by the gin automatically

	server.Run(":8080")

}

func getEvents(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{"meassage": "Hello!"})
	//gin.H is an alias of the map[string] any

}
