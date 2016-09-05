package interfaces

import "github.com/golang-es/go-cms/models"

type MetadataTypeDAO interface {
	InsertMetadataType(m *models.MetadataType) error
	UpdateMetadataType(m *models.MetadataType) error
	DeleteMetadataType(m *models.MetadataType) error
	GetByIDMetadataType(i int) (*models.MetadataType, error)
	GetAllMetadataType() ([]models.MetadataType, error)
}
