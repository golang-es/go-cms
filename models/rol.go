package models

import (
	"time"
)

// Rol modelo rol - perfil
type Rol struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
