package models

import (
	"time"
)

// RoleUser Modelo de la tabla pivote role_user
type RoleUser struct {
	ID        uint      `json:"id"`
	RoleID    uint      `json:"roleId"`
	UserID    uint      `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
