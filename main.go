package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	server.Run(":8080")
}
