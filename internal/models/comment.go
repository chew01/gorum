package models

import (
	"fmt"
	"gorum/internal/database"
	"strconv"
	"time"
)

type Comment struct {
	ID        int       `json:"id" example:"1"`
	Creator   User      `json:"creator"`
	Content   string    `json:"content" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit"`
	CreatedAt time.Time `json:"createdAt" example:"2023-01-01T00:00:00.000Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-01-01T00:00:00.000Z"`
}

type CommentRequest struct {
	Content string `json:"content" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit"`
}

func GetComment(db *database.Database, id string) (*Comment, error) {
	var commentData Comment
	stmt, err := db.Prepare("SELECT ID, creator, content, createdAt, updatedAt FROM comments WHERE ID = ?")
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()

	if err != nil {
		return nil, err
	}

	var createdAtStr string
	var updatedAtStr string
	res := stmt.QueryRow(id)
	if err = res.Scan(&commentData.ID, &commentData.Creator.Name, &commentData.Content,
		&createdAtStr, &updatedAtStr); err != nil {
		return nil, err
	}

	createdAt, err := time.Parse("2006-01-02T15:04:05-0700", createdAtStr)
	updatedAt, err := time.Parse("2006-01-02T15:04:05-0700", updatedAtStr)
	commentData.CreatedAt = createdAt
	commentData.UpdatedAt = updatedAt
	if err != nil {
		return nil, err
	}

	return &commentData, nil
}

func CreateComment(db *database.Database, comment *CommentRequest, postID string, creator string) (*Comment, error) {
	stmt, err := db.Prepare(`INSERT INTO comments (post, creator, content, createdAt, updatedAt)
									VALUES (?, ?, ?, ?, ?)`)
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()

	if err != nil {
		return nil, err
	}

	insert, err := stmt.Exec(postID, creator, comment.Content,
		time.Now().Format("2006-01-02T15:04:05-0700"), time.Now().Format("2006-01-02T15:04:05-0700"))
	if err != nil {
		return nil, err
	}

	lastInsertedID, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}

	commentData, err := GetComment(db, strconv.FormatInt(lastInsertedID, 10))

	return commentData, nil
}

func CheckUserComment(db *database.Database, id string) (*User, error) {
	var userData User
	stmt, err := db.Prepare("SELECT creator FROM comments WHERE id = ?")
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()

	if err != nil {
		return nil, err
	}

	res := stmt.QueryRow(id)
	if err := res.Scan(&userData.Name); err != nil {
		return nil, err
	}

	return &userData, nil
}

func UpdateComment(db *database.Database, comment *CommentRequest, id string) (*Comment, error) {
	stmt, err := db.Prepare(`UPDATE comments SET content = ?, updatedAt = ? WHERE ID = ? RETURNING ID`)
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()

	if err != nil {
		return nil, err
	}

	var returnID int
	res := stmt.QueryRow(comment.Content, time.Now().Format("2006-01-02T15:04:05-0700"), id)
	if err = res.Scan(&returnID); err != nil {
		return nil, err
	}

	commentData, err := GetComment(db, strconv.Itoa(returnID))
	if err != nil {
		return nil, err
	}

	return commentData, nil
}

func DeleteComment(db *database.Database, id string) (*Comment, error) {
	var commentData Comment
	stmt, err := db.Prepare(`DELETE FROM comments WHERE ID = ?
                  RETURNING ID, creator, content, createdAt, updatedAt`)
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()

	if err != nil {
		return nil, err
	}

	var createdAtStr string
	var updatedAtStr string
	res := stmt.QueryRow(id)
	if err = res.Scan(&commentData.ID, &commentData.Creator.Name, &commentData.Content,
		&createdAtStr, &updatedAtStr); err != nil {
		return nil, err
	}
	createdAt, err := time.Parse("2006-01-02T15:04:05-0700", createdAtStr)
	updatedAt, err := time.Parse("2006-01-02T15:04:05-0700", updatedAtStr)
	commentData.CreatedAt = createdAt
	commentData.UpdatedAt = updatedAt

	return &commentData, nil
}
