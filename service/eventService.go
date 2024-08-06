package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"project.com/event-booking/models"
	"project.com/event-booking/utils"
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

var uid = 0

func CreateEvent(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized! Empty  Token"})
		return
	}

	err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Failed !", "Error": err})
		return
	}

	var event models.Event
	uid++
	event.UserID = uid

	err = context.ShouldBindBodyWithJSON(&event) // this method maps the request body with the struct variablr event

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
	_, err = models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot fetch the event", "Error": err})
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

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No data with the given id", "Error": err})
		return
	}

	err = event.DeleteEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error occured in deletion", "Error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted Succesfully"})
}
