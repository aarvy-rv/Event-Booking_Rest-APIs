package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"project.com/event-booking/models"
)

func GetEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot pasre string to int64", "Error": err})
		return
	}
	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot fetch the event", "Error": err})
		return
	}
	context.JSON(http.StatusAccepted, event)

}

func GetEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Fetch Events", "Error": err})
		return
	}
	context.JSON(http.StatusOK, events)
	//gin.H is an alias of the map[string] any

}

func CreateEvent(context *gin.Context) {

	var event models.Event
	userId := context.GetInt64("userId")
	event.UserID = userId

	err := context.ShouldBindBodyWithJSON(&event) // this method maps the request body with the struct variablr event

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "required field missing"})
		return
	}

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Save Events", "Error": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created succesfully!"})

}

func UpdateEvent(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot pasre string to int64", "Error": err})
		return
	}
	event, err := models.GetEventByID(id)
	userId := context.GetInt64("userId")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot fetch the event", "Error": err})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event."})
		return

	}

	var updatedEvent models.Event

	err = context.ShouldBindBodyWithJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error", "Error": err})
		return
	}

	updatedEvent.Id = id
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error occured in update", "Error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated Succesfully"})

}

func Delete(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error occured in parsing string to int", "Error": err})
		return
	}

	event, err := models.GetEventByID(id)
	userId := context.GetInt64("userId")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No data with the given id", "Error": err})
		return
	}

	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to Delete this event."})
		return
	}

	err = event.DeleteEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error occured in deletion", "Error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted Succesfully"})
}
