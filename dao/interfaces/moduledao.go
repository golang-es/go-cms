package interfaces

import "github.com/golang-es/go-cms/models"

type ModuleDAO interface {
	InsertModule(m *models.Module) error
	UpdateModule(m *models.Module) error
	DeleteModule(m *models.Module) error
	GetByIDModule(i int) (*models.Module, error)
	GetAllModule() ([]models.Module, error)
}
