package models

import (
	"fmt"
	"gorum/internal/database"
	"strconv"
	"time"
)

type SimplePost struct {
	ID             int       `json:"id" example:"1"`
	Creator        User      `json:"creator"`
	Title          string    `json:"title" example:"The cat jumped over the lazy dog"`
	Synopsis       string    `json:"synopsis" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit"`
	CreatedAt      time.Time `json:"createdAt" example:"2023-01-01T00:00:00.000Z"`
	UpdatedAt      time.Time `json:"updatedAt" example:"2023-01-01T00:00:00.000Z"`
	LatestComments []Comment `json:"latestComments"`
	CommentCount   int       `json:"commentCount" example:"10"`
}

func GetSimplePosts(db *database.Database) ([]SimplePost, error) {
	var simplePosts []SimplePost
	stmt, err := db.Prepare(
		`SELECT
    			p.ID,
    			p.creator,
    			p.title,
    			substr(p.content, 1, 200) AS synopsis,
    			p.createdAt,
    			p.updatedAt,
    			(SELECT COUNT(*) FROM comments c WHERE c.post = p.id) AS commentCount
    			FROM posts AS p`)
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()

	stmtC, err := db.Prepare(`SELECT ID, creator, content, createdAt, updatedAt
					FROM comments WHERE post = ? ORDER BY createdAt DESC LIMIT 3`)
	defer func() {
		if err := stmtC.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := stmt.Query()
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Println("Error closing rows:", err)
			return
		}
	}()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp SimplePost
		var tempC []Comment
		var createdAtStr string
		var updatedAtStr string
		err = rows.Scan(&temp.ID, &temp.Creator.Name, &temp.Title, &temp.Synopsis,
			&createdAtStr, &updatedAtStr, &temp.CommentCount)
		if err != nil {
			return nil, err
		}

		createdAt, err := time.Parse("2006-01-02T15:04:05-0700", createdAtStr)
		updatedAt, err := time.Parse("2006-01-02T15:04:05-0700", updatedAtStr)
		temp.CreatedAt = createdAt
		temp.UpdatedAt = updatedAt
		if err != nil {
			return nil, err
		}

		rowsC, err := stmtC.Query(temp.ID)
		if err != nil {
			return nil, err
		}

		for rowsC.Next() {
			var temp Comment
			var createdAtStr string
			var updatedAtStr string
			err = rowsC.Scan(&temp.ID, &temp.Creator.Name, &temp.Content, &createdAtStr, &updatedAtStr)
			if err != nil {
				return nil, err
			}

			createdAt, err := time.Parse("2006-01-02T15:04:05-0700", createdAtStr)
			updatedAt, err := time.Parse("2006-01-02T15:04:05-0700", updatedAtStr)
			temp.CreatedAt = createdAt
			temp.UpdatedAt = updatedAt
			if err != nil {
				return nil, err
			}

			tempC = append(tempC, temp)
		}
		err = rowsC.Close()
		if err != nil {
			return nil, err
		}

		temp.LatestComments = tempC
		simplePosts = append(simplePosts, temp)
	}

	return simplePosts, nil
}

type Post struct {
	ID        int       `json:"id" example:"1"`
	Creator   User      `json:"creator"`
	Title     string    `json:"title" example:"The cat jumped over the lazy dog"`
	Content   string    `json:"content" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit"`
	CreatedAt time.Time `json:"createdAt" example:"2023-01-01T00:00:00.000Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-01-01T00:00:00.000Z"`
	Comments  []Comment `json:"comments"`
}

type PostRequest struct {
	Title   string `json:"title" example:"The cat jumped over the lazy dog"`
	Content string `json:"content" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit"`
}

func CreatePost(db *database.Database, post *PostRequest, creator string) (*Post, error) {
	var postData Post
	stmt, err := db.Prepare(`INSERT INTO posts (creator, title, content, createdAt, updatedAt)
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

	insert, err := stmt.Exec(creator, post.Title, post.Content,
		time.Now().Format("2006-01-02T15:04:05-0700"), time.Now().Format("2006-01-02T15:04:05-0700"))
	if err != nil {
		return nil, err
	}

	lastInsertedID, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}

	stmt, err = db.Prepare("SELECT ID, creator, title, content, createdAt, updatedAt FROM posts WHERE ID = ?")
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
	res := stmt.QueryRow(lastInsertedID)
	if err = res.Scan(&postData.ID, &postData.Creator.Name, &postData.Title,
		&postData.Content, &createdAtStr, &updatedAtStr); err != nil {
		return nil, err
	}

	createdAt, err := time.Parse("2006-01-02T15:04:05-0700", createdAtStr)
	updatedAt, err := time.Parse("2006-01-02T15:04:05-0700", updatedAtStr)
	postData.CreatedAt = createdAt
	postData.UpdatedAt = updatedAt
	if err != nil {
		return nil, err
	}

	return &postData, nil
}

func GetPost(db *database.Database, id string) (*Post, error) {
	var postData Post

	stmt, err := db.Prepare(
		`SELECT ID, creator, title, content, createdAt, updatedAt FROM posts WHERE ID = ?`,
	)
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Println("Error closing database statement:", err)
			return
		}
	}()
	stmtC, err := db.Prepare(
		`SELECT ID, post, creator, content, createdAt, updatedAt FROM comments WHERE post = ?`,
	)
	defer func() {
		if err := stmtC.Close(); err != nil {
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
	if err = res.Scan(&postData.ID, &postData.Creator.Name, &postData.Title,
		&postData.Content, &createdAtStr, &updatedAtStr); err != nil {
		return nil, err
	}
	createdAt, err := time.Parse("2006-01-02T15:04:05-0700", createdAtStr)
	updatedAt, err := time.Parse("2006-01-02T15:04:05-0700", updatedAtStr)
	postData.CreatedAt = createdAt
	postData.UpdatedAt = updatedAt
	if err != nil {
		return nil, err
	}

	rows, err := stmtC.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment Comment
		var createdAtStr string
		var updatedAtStr string

		err := rows.Scan(&comment.ID, &comment.Creator, &comment.Content, &createdAtStr, &updatedAtStr)
		if err != nil {
			return nil, err
		}
		createdAt, err := time.Parse("2006-01-02T15:04:05-0700", createdAtStr)
		updatedAt, err := time.Parse("2006-01-02T15:04:05-0700", updatedAtStr)
		comment.CreatedAt = createdAt
		comment.UpdatedAt = updatedAt

		postData.Comments = append(postData.Comments, comment)
	}

	return &postData, nil
}

func CheckUser(db *database.Database, id string) (*User, error) {
	var userData User
	stmt, err := db.Prepare("SELECT creator FROM posts WHERE id = ?")
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

func UpdatePost(db *database.Database, post *PostRequest, id string) (*Post, error) {
	stmt, err := db.Prepare(`UPDATE posts SET title = ?, content = ?, updatedAt = ?
             WHERE ID = ? RETURNING ID`)
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
	res := stmt.QueryRow(post.Title, post.Content, time.Now().Format("2006-01-02T15:04:05-0700"), id)
	if err = res.Scan(&returnID); err != nil {
		return nil, err
	}

	postData, err := GetPost(db, strconv.Itoa(returnID))
	if err != nil {
		return nil, err
	}

	return postData, nil
}

func DeletePost(db *database.Database, id string) (*Post, error) {
	var postData Post
	stmt, err := db.Prepare(`DELETE FROM posts WHERE ID = ?
                  RETURNING ID, creator, title, content, createdAt, updatedAt`)
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
	if err = res.Scan(&postData.ID, &postData.Creator.Name, &postData.Title,
		&postData.Content, &createdAtStr, &updatedAtStr); err != nil {
		return nil, err
	}
	createdAt, err := time.Parse("2006-01-02T15:04:05-0700", createdAtStr)
	updatedAt, err := time.Parse("2006-01-02T15:04:05-0700", updatedAtStr)
	postData.CreatedAt = createdAt
	postData.UpdatedAt = updatedAt

	return &postData, nil
}
