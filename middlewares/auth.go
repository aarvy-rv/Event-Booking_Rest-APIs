package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"project.com/event-booking/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized! Empty  Token"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization Failed !", "Error": err})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
