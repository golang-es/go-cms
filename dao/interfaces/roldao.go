package interfaces

import "github.com/golang-es/go-cms/models"

type RolDAO interface {
	InsertRol(r *models.Rol) error
	UpdateRol(r *models.Rol) error
	DeleteRol(r *models.Rol) error
	GetByIDRol(i uint) (*models.Rol, error)
	GetAllRol() ([]models.Rol, error)
}
