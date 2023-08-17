package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// ConnectDatabase establishes a connection to the SQLite database
func ConnectDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
