package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func CreateRepository() *Repository {
	userName := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	sslMode := "disable"
	query := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", userName, password, dbName, host, port, sslMode)
	db, err := sql.Open("postgres", query)
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{
		DB: db,
	}
}
