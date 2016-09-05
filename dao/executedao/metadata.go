package executedao

import "github.com/golang-es/go-cms/models"

// InsertMetadata Llama la función InsertMetadata del DAO
func InsertMetadata(m *models.Metadata) error {
	return metadataDAO.InsertMetadata(m)
}

// UpdateMetadata Llama la función UpdateMetadata del DAO
func UpdateMetadata(m *models.Metadata) error {
	return metadataDAO.InsertMetadata(m)
}

// DeleteMetadata llama la función DeleteMetadata del DAO
func DeleteMetadata(m *models.Metadata) error {
	return metadataDAO.DeleteMetadata(m)
}

// GetByIDMetadata llama la función GetByIDMetadata del DAO
func GetByIDMetadata(i int) (*models.Metadata, error) {
	return metadataDAO.GetByIDMetadata(i)
}

// GetAllMetadata llama la función GetAllmetadata del DAO
func GetAllMetadata() ([]models.Metadata, error) {
	return metadataDAO.GetAllMetadata()
}
