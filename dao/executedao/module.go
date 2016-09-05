package executedao

import "github.com/golang-es/go-cms/models"

// InsertModule llama la función insertar del dao
func InsertModule(m *models.Module) error {
	return moduleDAO.InsertModule(m)
}

// UpdateModule llama la función actualizar del dao
func UpdateModule(m *models.Module) error {
	return moduleDAO.UpdateModule(m)
}

// DeleteModule llama la función delete del dao
func DeleteModule(m *models.Module) error {
	return moduleDAO.DeleteModule(m)
}

// GetByIDModule llama la función getbyid del dao
func GetByIDModule(i int) (*models.Module, error) {
	return moduleDAO.GetByIDModule(i)
}

// GetAllModule llama la funcion getall del dao
func GetAllModule() ([]models.Module, error) {
	return moduleDAO.GetAllModule()
}
