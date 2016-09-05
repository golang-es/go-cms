package interfaces

import "github.com/golang-es/go-cms/models"

type PostDAO interface {
	InsertPost(p *models.Post) error
	UpdatePost(p *models.Post) error
	DeletePost(p *models.Post) error
	GetByIDPost(i int) (*models.Post, error)
	GetAllPost() ([]models.Post, error)
}
