package models

import "time"

type Comment struct {
	ID        int       `json:"id" example:"1"`
	Creator   User      `json:"creator"`
	Content   string    `json:"content" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit"`
	CreatedAt time.Time `json:"createdAt" example:"2023-01-01T00:00:00.000Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-01-01T00:00:00.000Z"`
}
