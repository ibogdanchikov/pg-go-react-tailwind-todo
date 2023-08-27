package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()

	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost",      // For running in Docker with Nginx.
		"http://localhost:3000", // For running with React.
	}
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "New phone, who dis?"})
	})

	return router
}
