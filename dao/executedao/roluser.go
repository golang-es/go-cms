package executedao

import "github.com/golang-es/go-cms/models"

// InsertRolUser llama la función insert del dao
func InsertRolUser(r *models.RoleUser) error {
	return rolUserDAO.InsertRolUser(r)
}

// UpdateRolUser llama la función update del dao
func UpdateRolUser(r *models.RoleUser) error {
	return rolUserDAO.UpdateRolUser(r)
}

// DeleteRolUser llama la función delete del dao
func DeleteRolUser(r *models.RoleUser) error {
	return rolUserDAO.DeleteRolUser(r)
}

// GetByIDRolUser llama la función getbyid del dao
func GetByIDRolUser(i int) (*models.RoleUser, error) {
	return rolUserDAO.GetByIDRolUser(i)
}

// GetAllRolUser llama la función getall del dao
func GetAllRolUser() ([]models.RoleUser, error) {
	return rolUserDAO.GetAllRolUser()
}
