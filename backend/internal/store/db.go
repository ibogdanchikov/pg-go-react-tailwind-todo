package store

import (
	"backend/internal/models"
	"database/sql"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(connStr string, migrationsDir string) {
	m, err := migrate.New(migrationsDir, connStr)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
		return
	}

	if err := m.Up(); err != nil {
		log.Printf("Migration Up: %v", err)
		return
	}
}

func InitDB(connStr string, maxRetries int) *sql.DB {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	for i := 0; i < maxRetries; i++ {
		err = db.Ping()

		if err == nil {
			return db
		}

		log.Printf("Failed to ping DB: %v, retrying...", err)
		time.Sleep(1 * time.Second)
	}

	log.Fatalf("Failed to ping DB after %d retries", maxRetries)
	return nil
}

func GetTasks(db *sql.DB) ([]models.Task, error) {
	rows, err := db.Query("SELECT id, description, done FROM task")
	if err != nil {
		log.Printf("Failed to query tasks: %v", err)
		return nil, err
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Done); err != nil {
			log.Printf("Failed to scan task: %v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Failed to iterate over rows: %v", err)
		return nil, err
	}

	return tasks, nil
}

func CreateTask(db *sql.DB, description string) (*models.Task, error) {

	row := db.QueryRow("INSERT INTO task (description) VALUES ($1) RETURNING id, description, done", description)

	var task models.Task
	err := row.Scan(&task.ID, &task.Description, &task.Done)

	if err != nil {
		log.Printf("Failed to scan inserted task: %v", err)
		return nil, err
	}

	return &task, nil
}
