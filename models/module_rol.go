package models

import (
    "errors"
    "time"

    conn "github.com/golang-es/go-cms/connection"
)

const (
    moduleByRolInsertPsql = "INSERT INTO module_rol (module_id, role_id, append, modify, remove, read) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at"
    moduleByRolUpdatePsql = "UPDATE module_rol SET append = $1, modify = $2, remove = $3, read = $4, updated_at = now() WHERE id = $5 RETURNING module_id, role_id, created_at, updated_at"
    moduleByRolDelete     = "DELETE FROM module_rol WHERE id = $1"
    moduleByRolSelectByID = "SELECT id, module_id, role_id, append, modify, remove, read, created_at, updated_at FROM module_rol WHERE id = $1"
    moduleByRolSelectAll  = "SELECT id, module_id, role_id, append, modify, remove, read, created_at, updated_at FROM module_rol"
)

// ModuleRol Modelo de la tabla pivote de modulos por rol
type ModuleRol struct {
    ID        uint
    ModuleID  uint
    RoleID    uint
    Append    bool
    Modify    bool
    Remove    bool
    Read      bool
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Create inserta un registro
func (mr *ModuleRol) Create() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = createModuleRolPostgresql(mr)
    case conn.MYSQL:
        err = errors.New("Not supported yet.")
    }
    return err
}

// Update actualiza el registro
func (mr *ModuleRol) Update() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = updateModuleRolPostgresql(mr)
    case conn.MYSQL:
        err = errors.New("Not supported yet.")
    }
    return err
}

// Delete borra un registro de la BD
func (mr *ModuleRol) Delete() error {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(moduleByRolDelete, mr.ID)
    if err != nil {
        return err
    }

    if i, _ := row.RowsAffected(); i == 1 {
        *mr = ModuleRol{}
    }
    return nil
}

// GetByID Obtiene un registro por el ID
func (mr *ModuleRol) GetByID() error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(moduleByRolSelectByID, mr.ID).Scan(&mr.ModuleID, &mr.RoleID, &mr.Append, &mr.Modify, &mr.Remove, &mr.Read, &mr.CreatedAt, &mr.UpdatedAt)
    return err
}

// GetAll Obtiene todos los registros
func (mr *ModuleRol) GetAll() ([]ModuleRol, error) {
    var modulesByRoles = make([]ModuleRol, 0)
    db := conn.Get()
    defer db.Close()

    rows, err := db.Query(moduleByRolSelectAll)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var moduleByRol ModuleRol
        err := rows.Scan(&moduleByRol.ID, &moduleByRol.ModuleID, &moduleByRol.RoleID, &moduleByRol.Append, &moduleByRol.Modify, &moduleByRol.Remove, &moduleByRol.Read, &moduleByRol.CreatedAt, &moduleByRol.UpdatedAt)
        if err != nil {
            return modulesByRoles, err
        }
        modulesByRoles = append(modulesByRoles, moduleByRol)
    }
    return modulesByRoles, nil
}

func createModuleRolPostgresql(mr *ModuleRol) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(moduleByRolInsertPsql, mr.ModuleID, mr.RoleID, mr.Append, mr.Modify, mr.Remove, mr.Read).Scan(&mr.ID, &mr.CreatedAt, &mr.UpdatedAt)
    return err
}

func updateModuleRolPostgresql(mr *ModuleRol) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(moduleByRolUpdatePsql, mr.Append, mr.Modify, mr.Remove, mr.Read, mr.ID).Scan(&mr.CreatedAt, &mr.UpdatedAt)
    return err
}
