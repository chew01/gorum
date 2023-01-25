package models

import (
	"fmt"
	"gorum/internal/database"
)

type User struct {
	Name string `json:"name" example:"John Doe"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

func SelectUser(db *database.Database, name string) (*User, error) {
	var userData User
	stmt, err := db.Prepare("SELECT name FROM users WHERE name = ?")
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()

	if err != nil {
		return nil, err
	}

	res := stmt.QueryRow(name)
	if err = res.Scan(&userData.Name); err != nil {
		return nil, err
	}

	return &userData, nil
}

func InsertUser(db *database.Database, user *User) (*User, error) {
	var userData User
	stmt, err := db.Prepare("INSERT INTO users (name) VALUES (?)")
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()

	if err != nil {
		return nil, err
	}

	insert, err := stmt.Exec(user.Name)
	if err != nil {
		return nil, err
	}

	lastInsertedID, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}

	stmt, err = db.Prepare("SELECT name FROM users WHERE ID = ?")
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()

	if err != nil {
		return nil, err
	}

	res := stmt.QueryRow(lastInsertedID)
	if err = res.Scan(&userData.Name); err != nil {
		return nil, err
	}

	return &userData, nil
}
