package models

import "time"

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

func GetAllPostSummaries() ([]SimplePost, error) {
	return []SimplePost{}, nil
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

func CreatePost() (Post, error) {
	return Post{}, nil
}
