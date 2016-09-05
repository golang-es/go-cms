package models

import (
	"time"
)

// MetadataType son los datos META que van en el HEAD del html
type MetadataType struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	MaxLength int       `json:"maxLength"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
