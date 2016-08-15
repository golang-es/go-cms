package models

import (
    "errors"
    "time"

    conn "github.com/golang-es/go-cms/connection"
)

const (
    metadataInsertPsql = "INSERT INTO metadatas (post_id, metadata_type_id, content) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
    metadataUpdatePsql = "UPDATE metadatas post_id = $1, metadata_type_id = $2, content = $3, updated_at = now() WHERE id = $4 RETURNING created_at, updated_at"
    metadataDelete     = "DELETE FROM metadatas WHERE id = $1"
    metadataSelectByID = "SELECT id, post_id, metadata_type_id, content, created_at, updated_at FROM metadatas WHERE id = $1"
    metadataSelectAll  = "SELECT id, post_id, metadata_type_id, content, created_at, updated_at FROM metadatas"
)

// Metadata es el contenido que irá en el HEAD
type Metadata struct {
    ID             int
    PostID         int
    MetadataTypeID int
    Content        string
    CreatedAt      time.Time
    UpdatedAt      time.Time
}

// Create crea un registro en la BD
func (m *Metadata) Create() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = createMetadataPostgresql(m)
    case conn.MYSQL:
        err = errors.New("Not supported yet.")
    }
    return err
}

// Update actualiza un registro en la BD
func (m *Metadata) Update() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = updateMetadataPostgresql(m)
    case conn.MYSQL:
        err = errors.New("Not supported yet.")
    }
    return err
}

// Delete borra un registro de la BD
func (m *Metadata) Delete() error {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(metadataDelete, m.ID)
    if err != nil {
        return err
    }

    if i, _ := row.RowsAffected(); i == 1 {
        *m = Metadata{}
    }
    return nil
}

// GetByID obtiene un registro de la BD
func (m *Metadata) GetByID() error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(metadataSelectByID, m.ID).Scan(&m.ID, &m.PostID, &m.MetadataTypeID, &m.Content, &m.CreatedAt, &m.UpdatedAt)
    return err
}

// GetAll obtiene todos los registros de la BD
func (m *Metadata) GetAll() ([]Metadata, error) {
    var metadatas []Metadata
    db := conn.Get()
    defer db.Close()

    rows, err := db.Query(metadataSelectAll)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var md Metadata
        err := rows.Scan(&md.ID, &md.PostID, &md.MetadataTypeID, &md.Content, &md.CreatedAt, &md.UpdatedAt)
        if err != nil {
            return metadatas, err
        }
        metadatas = append(metadatas, md)
    }
    return metadatas, nil
}

// Crea un registro en Postgresql
func createMetadataPostgresql(m *Metadata) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(metadataInsertPsql, m.PostID, m.MetadataTypeID, m.Content).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
    return err
}

func updateMetadataPostgresql(m *Metadata) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(metadataUpdatePsql, m.MetadataTypeID, m.Content, m.PostID).Scan(&m.CreatedAt, &m.UpdatedAt)
    return err
}
