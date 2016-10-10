package executedao

import "github.com/golang-es/go-cms/models"

// InsertUser llama la función insert del dao
func InsertUser(u *models.User) error {
	return userDAO.InsertUser(u)
}

// UpdateUser llama la función update del dao
func UpdateUser(u *models.User) error {
	return userDAO.UpdateUser(u)
}

// DeleteUser llama la función delete del dao
func DeleteUser(u *models.User) error {
	return userDAO.DeleteUser(u)
}

// GetByIDUser llama la función getbyid del dao
func GetByIDUser(i int) (*models.User, error) {
	return userDAO.GetByIDUser(i)
}

// GetAllUser llama la función getalluser del dao
func GetAllUser() ([]models.User, error) {
	return userDAO.GetAllUser()
}

// GetLogin llama la función getlogin del dao
func GetLogin(e string, p string) (*models.User, error) {
	return userDAO.GetLogin(e, p)
}

// GetUserByEmail llama la función getuserbyemail del dao
func GetUserByEmail(e string) (*models.User, error) {
	return userDAO.GetUserByEmail(e)
}
