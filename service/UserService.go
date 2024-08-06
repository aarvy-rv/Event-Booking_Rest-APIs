package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"project.com/event-booking/models"
	"project.com/event-booking/utils"
)

func Signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error in mapping request body and user model", "Error": err})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error in saving user", "Error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User successfully signed up!"})

}

func Login(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error in mapping request body and user model", "Error": err})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user", "Error": err})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not generate token", "Error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User Authenticated!", "token": token})

}
