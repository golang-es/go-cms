package models

import (
	"time"
)

// Post post que publican los usuarios del cms
type Post struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"userId"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"publishedAt"`
	Poster      string    `json:"poster"`
	Banner      string    `json:"banner"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
