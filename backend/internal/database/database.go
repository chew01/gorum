package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type Database struct {
	*sql.DB
}

func New(dbPath string) (*Database, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%v?_foreign_keys=on", dbPath))
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (db *Database) Init() error {
	file, err := os.ReadFile("internal/database/init.sql")
	_, err = db.Exec(string(file))
	return err
}
