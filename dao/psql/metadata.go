package psql

import "github.com/golang-es/go-cms/models"

type MetadataDAOPSQL struct{}

// InsertMetadata Inserta un registro en la BD
func (m MetadataDAOPSQL) InsertMetadata(metadata *models.Metadata) error {
	query := "INSERT INTO metadatas (post_id, metadata_type_id, content) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, metadata.PostID, metadata.MetadataTypeID, metadata.Content).Scan(&metadata.ID, &metadata.CreatedAt, &metadata.UpdatedAt)
	return err
}

// Updatemetadata actualiza un registro en la BD
func (m MetadataDAOPSQL) UpdateMetadata(metadata *models.Metadata) error {
	query := "UPDATE metadatas post_id = $1, metadata_type_id = $2, content = $3, updated_at = now() WHERE id = $4 RETURNING created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, metadata.MetadataTypeID, metadata.Content, metadata.PostID).Scan(&metadata.CreatedAt, &metadata.UpdatedAt)
	return err
}

// DeleteMetadata borra un registro en la BD
func (m MetadataDAOPSQL) DeleteMetadata(metadata *models.Metadata) error {
	query := "DELETE FROM metadatas WHERE id = $1"
	db := get()
	defer db.Close()

	row, err := db.Exec(query, metadata.ID)
	if err != nil {
		return err
	}

	if i, _ := row.RowsAffected(); i == 1 {
		metadata = &models.Metadata{}
	}
	return nil
}

// GetByIDMetadata obtiene un registro en la BD
func (m MetadataDAOPSQL) GetByIDMetadata(id int) (*models.Metadata, error) {
	query := "SELECT id, post_id, metadata_type_id, content, created_at, updated_at FROM metadatas WHERE id = $1"
	metadata := &models.Metadata{}
	db := get()
	defer db.Close()

	err := db.QueryRow(query, id).Scan(&metadata.ID, &metadata.PostID, &metadata.MetadataTypeID, &metadata.Content, &metadata.CreatedAt, &metadata.UpdatedAt)
	return metadata, err

}

// GetAllMetadata obtiene todos los registros de la BD
func (m MetadataDAOPSQL) GetAllMetadata() ([]models.Metadata, error) {
	query := "SELECT id, post_id, metadata_type_id, content, created_at, updated_at FROM metadatas"
	var metadatas []models.Metadata
	db := get()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var md models.Metadata
		err := rows.Scan(&md.ID, &md.PostID, &md.MetadataTypeID, &md.Content, &md.CreatedAt, &md.UpdatedAt)
		if err != nil {
			return metadatas, err
		}
		metadatas = append(metadatas, md)
	}
	return metadatas, nil
}
