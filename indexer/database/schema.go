package database

import (
	"database/sql"
)

// InitSchema initializes the database schema
func InitSchema(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS erc20_transfers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		block_number INTEGER NOT NULL,
		tx_hash TEXT NOT NULL,
		from_address TEXT NOT NULL,
		to_address TEXT NOT NULL,
		value TEXT NOT NULL
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
