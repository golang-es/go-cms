package models

import (
    "errors"
    "time"

    conn "github.com/golang-es/go-cms/connection"
)

const (
    roleByUserInsertPsql = "INSERT INTO role_user (role_id, user_id) VALUES ($1, $1) RETURNING id, created_at, updated_at"
    roleByUserUpdatePsql = "UPDATE role_user SET role_id = $1, user_id = $2, updated_at = now() WHERE id = $3 RETURNING created_at, updated_at"
    roleByUserDelete     = "DELETE FROM role_user WHERE id = $1"
    roleByUserSelectByID = "SELECT id, role_id, user_id, created_at, updated_at FROM role_user WHERE id = $1"
    roleByUserSelectAll  = "SELECT id, role_id, user_id, created_at, updated_at FROM role_user"
)

// RoleUser Modelo de la tabla pivote role_user
type RoleUser struct {
    ID        uint
    RoleID    uint
    UserID    uint
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Create crea un registro en la BD
func (ru *RoleUser) Create() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = createRoleByUserPostgresql(ru)
    case conn.MYSQL:
        err = errors.New("Not supported yet.")
    }
    return err
}

// Update actualiza un registro en la BD
func (ru *RoleUser) Update() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = updateRoleByUserPostgresql(ru)
    case conn.MYSQL:
        err = errors.New("Not supported yet.")
    }
    return err
}

// Delete borra un registro en la BD
func (ru *RoleUser) Delete() error {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(roleByUserDelete, ru.ID)
    if err != nil {
        return err
    }

    if i, _ := row.RowsAffected(); i == 1 {
        *ru = RoleUser{}
    }
    return nil
}

// GetByID Obtiene un registro por el ID
func (ru *RoleUser) GetByID() error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(roleByUserSelectByID, ru.ID).Scan(&ru.ID, &ru.RoleID, &ru.UserID, &ru.CreatedAt, &ru.UpdatedAt)
    return err
}

// GetAll Obtiene todos los registros
func (ru *RoleUser) GetAll() ([]RoleUser, error) {
    var roleByUsers = make([]RoleUser, 0)
    db := conn.Get()
    defer db.Close()

    rows, err := db.Query(roleByUserSelectAll)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var roleByUser RoleUser
        err := rows.Scan(&roleByUser.ID, &roleByUser.RoleID, &roleByUser.UserID, &roleByUser.CreatedAt, &roleByUser.UpdatedAt)
        if err != nil {
            return roleByUsers, err
        }
        roleByUsers = append(roleByUsers, roleByUser)
    }
    return roleByUsers, nil
}

func createRoleByUserPostgresql(ru *RoleUser) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(roleByUserInsertPsql, ru.RoleID, ru.UserID).Scan(&ru.ID, &ru.CreatedAt, &ru.UpdatedAt)
    return err
}

func updateRoleByUserPostgresql(ru *RoleUser) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(roleByUserUpdatePsql, ru.RoleID, ru.UserID, ru.ID).Scan(&ru.CreatedAt, &ru.UpdatedAt)
    return err
}
