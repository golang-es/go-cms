package models

import (
    "time"

    conn "github.com/golang-es/go-cms/connection"
)

const (
    rolInsertPsql  = "INSERT INTO roles (name) VALUES ($1) RETURNING id, active, created_at, updated_at"
    rolInsertMysql = "INSERT INTO roles (name) VALUES ($1)"
    rolUpdatePsql  = "UPDATE roles SET name = $1, active = $2, updated_at = now() WHERE id = $3  RETURNING created_at, updated_at"
    rolUpdateMysql = "UPDATE roles SET name = $1, active = $2, updated_at = now() WHERE id = $3"
    rolDelete      = "DELETE FROM roles WHERE id = $1"
    rolSelectByID  = "SELECT id, name, active, created_at, updated_at FROM roles WHERE id = $1"
    rolSelectAll   = "SELECT id, name, active, created_at, updated_at FROM roles"
)

// Rol modelo rol - perfil
type Rol struct {
    ID        uint
    Name      string
    Active    bool
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Create Inserta un registro en la BD
func (r *Rol) Create() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = createRolPostgresql(r)
    case conn.MYSQL:
        err = createRolMysql(r)
    }
    return err
}

// Update actualiza un registro en la BD
func (r *Rol) Update() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = updateRolPostgresql(r)
    case conn.MYSQL:
        err = updateRolMysql(r)
    }
    return err
}

// Delete borra un registro de la BD
func (r *Rol) Delete() error {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(rolDelete, r.ID)
    if err != nil {
        return err
    }

    if i, _ := row.RowsAffected(); i == 1 {
        *r = Rol{}
    }
    return nil
}

// GetByID Obtiene un rol por su ID
func (r *Rol) GetByID() error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(rolSelectByID, r.ID).Scan(&r.ID, &r.Name, &r.Active, &r.CreatedAt, &r.UpdatedAt)
    return err
}

// GetAll Obtiene todos los registros de la tabla
func (r *Rol) GetAll() ([]Rol, error) {
    var roles = make([]Rol, 0)
    db := conn.Get()
    defer db.Close()

    rows, err := db.Query(rolSelectAll)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var rol Rol
        err := rows.Scan(&rol.ID, &rol.Name, &rol.Active, &rol.CreatedAt, &rol.UpdatedAt)
        if err != nil {
            return roles, err
        }
        roles = append(roles, rol)
    }
    return roles, nil
}

func createRolPostgresql(r *Rol) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(rolInsertPsql, r.Name).Scan(&r.ID, &r.Active, &r.CreatedAt, &r.UpdatedAt)
    return err
}

func updateRolPostgresql(r *Rol) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(rolUpdatePsql, r.Name, r.Active, r.ID).Scan(&r.CreatedAt, &r.UpdatedAt)
    return err
}

func createRolMysql(r *Rol) error {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(rolInsertMysql, r.Name)
    if err != nil {
        return err
    }
    id, err := row.LastInsertId()
    if err != nil {
        return err
    }
    r.ID = uint(id)
    r.GetByID()
    return nil
}

func updateRolMysql(r *Rol) error {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(rolUpdateMysql, r.Name, r.Active, r.ID)
    if err != nil {
        return err
    }
    _, err = row.RowsAffected()
    if err != nil {
        return err
    }
    r.GetByID()
    return nil
}
