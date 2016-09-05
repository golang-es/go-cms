package executedao

import "github.com/golang-es/go-cms/models"

// InsertModuleRol llama la función insert del dao
func InsertModuleRol(m *models.ModuleRol) error {
	return moduleRolDAO.InsertModuleRol(m)
}

// UpdateModuleRol llama la función update del dao
func UpdateModuleRol(m *models.ModuleRol) error {
	return moduleRolDAO.UpdateModuleRol(m)
}

// DeleteModuleRol llama la función delete del dao
func DeleteModuleRol(m *models.ModuleRol) error {
	return moduleRolDAO.DeleteModuleRol(m)
}

// GetByIDModuleRol llama la función getbyid del dao
func GetByIDModuleRol(i int) (*models.ModuleRol, error) {
	return moduleRolDAO.GetByIDModuleRol(i)
}

// GetAllModuleRol llama la función getall del dao
func GetAllModuleRol() ([]models.ModuleRol, error) {
	return moduleRolDAO.GetAllModuleRol()
}
