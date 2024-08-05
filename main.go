package main

import (
	"github.com/gin-gonic/gin"
	"project.com/event-booking/db"
	"project.com/event-booking/handlers"
)

func main() {

	db.InitDB()
	server := gin.Default()
	handlers.Apis(server)

}
