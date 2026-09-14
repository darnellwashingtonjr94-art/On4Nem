package storage

import (
	"database/sql"
	"log"
)

func InitializeDBPool(connectionString string) (*sql.DB, error) {
	log.Println("Initializing high-availability PostgreSQL connection pool for On4Nem analytics...")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	return db, nil
}
