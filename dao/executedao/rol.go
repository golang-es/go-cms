package executedao

import "github.com/golang-es/go-cms/models"

// InsertRol llama la función InsertRol del DAO
func InsertRol(r *models.Rol) error {
	return rolDAO.InsertRol(r)
}

// UpdateRol llama la función UpdateRol del DAO
func UpdateRol(r *models.Rol) error {
	return rolDAO.UpdateRol(r)
}

// DeleteRol Llama la función DeleteRol del DAO
func DeleteRol(r *models.Rol) error {
	return rolDAO.DeleteRol(r)
}

// GetByIDRol Llama la función GetByIDRol del DAO
func GetByIDRol(id uint) (*models.Rol, error) {
	return rolDAO.GetByIDRol(id)
}

// GetAllRol Llama la función GetAllRol del DAO
func GetAllRol() ([]models.Rol, error) {
	return rolDAO.GetAllRol()
}
