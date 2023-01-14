package database

import (
	"database/sql"
	"errors"
)

const path string = "root.db"

type Database struct {
	client *sql.DB
}

func Get() (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, errors.New("failed to open sql from given path")
	}
	return &Database{client: db}, nil
}

func Init() error {
	conn, err := Get()
	if err != nil {
		return err
	}

	return nil
}
