package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// app := gin.Default()
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	log.Fatal(app.Run(":8080"))
}
