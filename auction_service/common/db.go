package dbaccessor

import (
	"database/sql"
	"fmt"
	"os"
)

type Repository struct {
	DB *sql.DB
}

func InitDB() (*Repository, error) {
	fmt.Printf("Init Data Source\n")

	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDB := os.Getenv("PG_DB")
	pgSSL := "disable"

	pgConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", pgHost, pgPort, pgUser, pgPassword, pgDB, pgSSL)

	db, err := sql.Open("postgres", pgConnString)
  if err != nil {
    return nil, fmt.Errorf("error opening db: %w", err)
  }
 
	// Verify database connection is working
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	return &Repository{
		DB: db }, nil
}

func (d *Repository) close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing Postgresql: %w", err)
	}
	
	return nil
}