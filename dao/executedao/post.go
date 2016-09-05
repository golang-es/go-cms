package executedao

import "github.com/golang-es/go-cms/models"

// InsertPost llama la función insert del dao
func InsertPost(p *models.Post) error {
	return postDAO.InsertPost(p)
}

// UpdatePost llama la función update del dao
func UpdatePost(p *models.Post) error {
	return postDAO.UpdatePost(p)
}

// DeletePost llama la función delete del dao
func DeletePost(p *models.Post) error {
	return postDAO.DeletePost(p)
}

// GetByIDPost llama la función getbyid del dao
func GetByIDPost(i int) (*models.Post, error) {
	return postDAO.GetByIDPost(i)
}

// GetAllPost llama la función getall del dao
func GetAllPost() ([]models.Post, error) {
	return postDAO.GetAllPost()
}
