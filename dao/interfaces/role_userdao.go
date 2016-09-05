package interfaces

import "github.com/golang-es/go-cms/models"

type RolUserDAO interface {
	InsertRolUser(r *models.RoleUser) error
	UpdateRolUser(r *models.RoleUser) error
	DeleteRolUser(r *models.RoleUser) error
	GetByIDRolUser(i int) (*models.RoleUser, error)
	GetAllRolUser() ([]models.RoleUser, error)
}
