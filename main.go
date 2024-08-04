package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"project.com/event-booking/models"
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

	server.POST("/create/event", createEvent)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {

	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
	//gin.H is an alias of the map[string] any

}

var id = 0
var uid = 0

func createEvent(context *gin.Context) {

	var event models.Event
	id++
	uid++
	event.Id = id
	event.UserID = uid

	err := context.ShouldBindBodyWithJSON(&event) // this method maps the request body with the struct variablr event

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "required field missing"})
		return
	}

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created succesfully!"})

}
