package dbConn

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"os"
)

type DB struct {
	*sql.DB
}

func Init() (*DB, error) {

	cfg, err := NewDBConfig()
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(cfg.Path)

	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to check db file: %w", err)
	}

	db, err := sql.Open("sqlite", cfg.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	database := &DB{db}

	if err := database.migrate(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("failed to migrate db: %w", err)
	}
	return database, nil
}

func (d *DB) migrate() error {
	const query = `
		CREATE TABLE IF NOT EXISTS scheduler (
			id       INTEGER PRIMARY KEY AUTOINCREMENT,
			date     TEXT NOT NULL,
			title    TEXT NOT NULL,
			comment  TEXT,
			repeat   TEXT
		);
	`

	_, err := d.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
