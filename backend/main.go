package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	runMigrations()

	router := setupRouter()

	router.Run(":8080")
}

func runMigrations() {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://todo:todo@db:5432/todo?sslmode=disable",
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
		return
	}

	if err := m.Up(); err != nil {
		log.Printf("Migration Up: %v", err)
		return
	}
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
