package scripts

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) error {
	log.Println("Executing database migrations for On4Nem tables (trades, ledgers, audit logs)...")
	query := `
	CREATE TABLE IF NOT EXISTS trades (
		id SERIAL PRIMARY KEY,
		market_id VARCHAR(255),
		stake NUMERIC,
		odds NUMERIC,
		status VARCHAR(50),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}
