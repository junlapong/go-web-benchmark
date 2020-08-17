package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	gin.SetMode(gin.ReleaseMode)

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

	log.Fatal(app.Run(":" + port))
}
