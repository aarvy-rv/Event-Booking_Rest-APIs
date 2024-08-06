package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"project.com/event-booking/models"
)

func RegisterForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot pasre string to int64", "Error": err})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered Successfully!"})

}

func CancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cannot pasre string to int64", "Error": err})
		return
	}

	var event models.Event
	event.Id = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Canceled Successfully!"})

}
