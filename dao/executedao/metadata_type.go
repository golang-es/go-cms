package executedao

import "github.com/golang-es/go-cms/models"

// InsertMetadataType llama la función insertar del dao
func InsertMetadataType(m *models.MetadataType) error {
	return metadataTypeDAO.InsertMetadataType(m)
}

// UpdatemetadataType llama la función actualizar del dao
func UpdateMetadataType(m *models.MetadataType) error {
	return metadataTypeDAO.UpdateMetadataType(m)
}

// DeleteMetadataType llama la función borrar del dao
func DeleteMetadataType(m *models.MetadataType) error {
	return metadataTypeDAO.DeleteMetadataType(m)
}

// GetByIDMetadataType llama la función obtener por id del dao
func GetByIDMetadataType(i int) (*models.MetadataType, error) {
	return metadataTypeDAO.GetByIDMetadataType(i)
}

// GetAllMetadataType llama la función obtener todos del dao
func GetAllMetadataType() ([]models.MetadataType, error) {
	return metadataTypeDAO.GetAllMetadataType()
}
