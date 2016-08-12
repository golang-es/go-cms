package models

import (
    "errors"
    "time"

    conn "github.com/golang-es/go-cms/connection"
)

const (
    moduleInsertPsql = "INSERT INTO modules (name, description) VALUES ($1, $2) RETURNING id, created_at, updated_at"
    moduleUpdatePsql = "UPDATE modules SET name = $1, description = $2, updated_at = now() WHERE id = $3  RETURNING created_at, updated_at"
    moduleDelete     = "DELETE FROM modules WHERE id = $1"
    moduleSelectByID = "SELECT id, name, description, created_at, updated_at FROM modules WHERE id = $1"
    moduleSelectAll  = "SELECT id, name, description, created_at, updated_at FROM roles"
)

// Module módulo del cms
type Module struct {
    ID          uint
    Name        string
    Description string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// Create inserta un nuevo registro
func (m *Module) Create() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = createModulePostgresql(m)
    case conn.MYSQL:
        err = errors.New("Not supported yet.")
    }
    return err
}

// Update actualiza el registro
func (m *Module) Update() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = updateModulePostgresql(m)
    case conn.MYSQL:
        err = errors.New("Not supported yet.")
    }
    return err
}

// Delete Borra un registro de la BD
func (m *Module) Delete() error {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(moduleDelete, m.ID)
    if err != nil {
        return err
    }

    if i, _ := row.RowsAffected(); i == 1 {
        *m = Module{}
    }
    return nil
}

// GetByID Obtiene un registro por el ID
func (m *Module) GetByID() error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(moduleSelectByID, m.ID).Scan(&m.ID, &m.Name, &m.Description, &m.CreatedAt, &m.UpdatedAt)
    return err
}

// GetAll obtiene todos los registros de la BD
func (m *Module) GetAll() ([]Module, error) {
    var modules = make([]Module, 0)
    db := conn.Get()
    defer db.Close()

    rows, err := db.Query(moduleSelectAll)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var module Module
        err := rows.Scan(&module.ID, &module.Name, &module.Description, &module.CreatedAt, &module.UpdatedAt)
        if err != nil {
            return modules, err
        }
        modules = append(modules, module)
    }
    return modules, nil
}

func createModulePostgresql(m *Module) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(moduleInsertPsql, m.Name, m.Description).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
    return err
}

func updateModulePostgresql(m *Module) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(moduleUpdatePsql, m.Name, m.Description, m.ID).Scan(&m.CreatedAt, &m.UpdatedAt)
    return err
}
