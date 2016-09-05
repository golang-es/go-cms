package models

import (
	"time"
)

// ModuleRol Modelo de la tabla pivote de modulos por rol
type ModuleRol struct {
	ID        uint      `json:"id"`
	ModuleID  uint      `json:"moduleId"`
	RoleID    uint      `json:"roleId"`
	Append    bool      `json:"append"`
	Modify    bool      `json:"modify"`
	Remove    bool      `json:"remove"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
