package models

import (
	"time"
)

// Metadata es el contenido que irá en el HEAD
type Metadata struct {
	ID             int `json:"id"`
	PostID         int `json:"postId"`
	MetadataTypeID int `json:"metadataTypeId"`
	Content        string `json:"content"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
