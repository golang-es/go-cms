package interfaces

import "github.com/golang-es/go-cms/models"

type ModuleRolDAO interface {
	InsertModuleRol(m *models.ModuleRol) error
	UpdateModuleRol(m *models.ModuleRol) error
	DeleteModuleRol(m *models.ModuleRol) error
	GetByIDModuleRol(i int) (*models.ModuleRol, error)
	GetAllModuleRol() ([]models.ModuleRol, error)
}
