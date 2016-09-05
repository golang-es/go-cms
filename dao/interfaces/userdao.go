package interfaces

import "github.com/golang-es/go-cms/models"

type UserDAO interface {
	InsertUser(u *models.User) error
	UpdateUser(u *models.User) error
	DeleteUser(u *models.User) error
	GetByIDUser(i int) (*models.User, error)
	GetAllUser() ([]models.User, error)
}
