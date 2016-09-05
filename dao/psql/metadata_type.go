package psql

import "github.com/golang-es/go-cms/models"

type MetadataTypeDAOPSQL struct{}

// InsertMetadataType inserta un registro en la BD
func (m MetadataTypeDAOPSQL) InsertMetadataType(meta *models.MetadataType) error {
	query := "INSERT INTO metadata_types (type, max_length) VALUES ($1, $2) RETURNING id, created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, meta.ID, meta.Type, meta.MaxLength).Scan(&meta.ID, &meta.CreatedAt, &meta.UpdatedAt)
	return err
}

// UpdateMetadataType actualiza un registro en la BD
func (m MetadataTypeDAOPSQL) UpdateMetadataType(meta *models.MetadataType) error {
	query := "UPDATE metadata_types SET type = $1, max_length = $2, updated_at = now() WHERE id = $3 RETURNING created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, meta.Type, meta.MaxLength, meta.ID).Scan(&meta.CreatedAt, &meta.UpdatedAt)
	return err
}

// DeleteMetadataType borra un registro en la BD
func (m MetadataTypeDAOPSQL) DeleteMetadataType(meta *models.MetadataType) error {
	query := "DELETE FROM metadata_types WHERE id = $1"
	db := get()
	defer db.Close()

	row, err := db.Exec(query, meta.ID)
	if err != nil {
		return err
	}

	if i, _ := row.RowsAffected(); i == 1 {
		meta = &models.MetadataType{}
	}
	return nil
}

// GetByIDMetadataType consulta un registro en la BD
func (m MetadataTypeDAOPSQL) GetByIDMetadataType(id int) (*models.MetadataType, error) {
	query := "SELECT id, type, max_length, created_at, updated_at FROM metadata_types WHERE ID = $1"
	meta := &models.MetadataType{}
	db := get()
	defer db.Close()

	err := db.QueryRow(query, id).Scan(&meta.ID, &meta.Type, &meta.MaxLength, &meta.CreatedAt, &meta.UpdatedAt)
	return meta, err
}

// GetAllMetadataType consulta todos los registros en la BD
func (m MetadataTypeDAOPSQL) GetAllMetadataType() ([]models.MetadataType, error) {
	query := "SELECT id, type, max_length, created_at, updated_at FROM metadata_types"
	var metadatatypes []models.MetadataType
	db := get()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mt models.MetadataType
		err := rows.Scan(&mt.ID, &mt.Type, &mt.MaxLength, &mt.CreatedAt, &mt.UpdatedAt)
		if err != nil {
			return metadatatypes, err
		}
		metadatatypes = append(metadatatypes, mt)
	}
	return metadatatypes, nil
}
