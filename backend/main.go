package main

import (
	"backend/api/handlers"
	"backend/internal/store"
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	connStr := "postgres://todo:todo@db:5434/todo?sslmode=disable"

	store.RunMigrations(connStr, "file://db/migrations")

	db := store.InitDB(connStr, 10)
	defer db.Close()

	router := setupRouter(db)
	router.Run(":8080")
}

func setupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost",      // For running in Docker with Nginx.
		"http://localhost:3000", // For running with React.
	}
	router.Use(cors.New(config))

	handlers.GetTasks(router, db)
	handlers.CreateTask(router, db)

	return router
}
