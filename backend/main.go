package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()

	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "New phone, who dis?"})
	})

	return router
}
