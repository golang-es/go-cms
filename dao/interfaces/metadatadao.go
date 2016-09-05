package interfaces

import "github.com/golang-es/go-cms/models"

type MetadataDAO interface {
	InsertMetadata(m *models.Metadata) error
	UpdateMetadata(m *models.Metadata) error
	DeleteMetadata(m *models.Metadata) error
	GetByIDMetadata(i int) (*models.Metadata, error)
	GetAllMetadata() ([]models.Metadata, error)
}
