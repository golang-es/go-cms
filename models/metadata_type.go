package models

import (
	"errors"
	"time"

	conn "github.com/golang-es/go-cms/connection"
)

const (
	metadataTypeInsertPsql = "INSERT INTO metadata_types (type, max_length) VALUES ($1, $2) RETURNING id, created_at, updated_at"
	metadataTypeUpdatePsql = "UPDATE metadata_types SET type = $1, max_length = $2, updated_at = now() WHERE id = $3 RETURNING created_at, updated_at"
	metadataTypeDelete     = "DELETE FROM metadata_types WHERE id = $1"
	metadataTypeSelectByID = "SELECT id, type, max_length, created_at, updated_at FROM metadata_types WHERE ID = $1"
	metadataTypeSelectAll  = "SELECT id, type, max_length, created_at, updated_at FROM metadata_types"
)

// MetadataType son los datos META que van en el HEAD del html
type MetadataType struct {
	ID        int
	Type      string
	MaxLength int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create inserta un registro en la BD
func (m *MetadataType) Create() error {
	var err error
	switch conn.GetNameEngine() {
	case conn.POSTGRESQL:
		err = createMetadataTypePostgresql(m)
	case conn.MYSQL:
		err = errors.New("Not supported yet")
	}
	return err
}

// Update actualiza un registro en la BD
func (m *MetadataType) Update() error {
	var err error
	switch conn.GetNameEngine() {
	case conn.POSTGRESQL:
		err = updateMetadataTypePostgresql(m)
	case conn.MYSQL:
		err = errors.New("Not supported yet")
	}
	return err
}

// Delete Borra un registro de la BD
func (m *MetadataType) Delete() error {
	db := conn.Get()
	defer db.Close()

	row, err := db.Exec(metadataTypeDelete, m.ID)
	if err != nil {
		return err
	}

	if i, _ := row.RowsAffected(); i == 1 {
		*m = MetadataType{}
	}
	return nil
}

// GetByID obtiene un registro de la BD
func (m *MetadataType) GetByID() error {
	db := conn.Get()
	defer db.Close()

	err := db.QueryRow(metadataTypeSelectByID, m.ID).Scan(&m.ID, &m.Type, &m.MaxLength, &m.CreatedAt, &m.UpdatedAt)
	return err
}

// GetAll obtiene todos los registros de la BD
func (m *MetadataType) GetAll() ([]MetadataType, error) {
	var metadatatypes []MetadataType
	db := conn.Get()
	defer db.Close()

	rows, err := db.Query(metadataTypeSelectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mt MetadataType
		err := rows.Scan(&mt.ID, &mt.Type, &mt.MaxLength, &mt.CreatedAt, &mt.UpdatedAt)
		if err != nil {
			return metadatatypes, err
		}
		metadatatypes = append(metadatatypes, mt)
	}
	return metadatatypes, nil
}

// Crea un registro en la bd de POSTGRESQL
func createMetadataTypePostgresql(m *MetadataType) error {
	db := conn.Get()
	defer db.Close()

	err := db.QueryRow(metadataTypeInsertPsql, m.ID, m.Type, m.MaxLength).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	return err
}

// Actualiza un registro en la bd de POSTGRESQL
func updateMetadataTypePostgresql(m *MetadataType) error {
	db := conn.Get()
	defer db.Close()

	err := db.QueryRow(metadataTypeUpdatePsql, m.Type, m.MaxLength, m.ID).Scan(&m.CreatedAt, &m.UpdatedAt)
	return err
}
