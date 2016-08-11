package models

import (
    "database/sql"
    "log"
    "time"

    conn "github.com/golang-es/go-cms/connection"
)

const (
    insertPsql = "INSERT INTO roles (name) VALUES ($1) RETURNING id, name, active, created_at, updated_at"
    insertMsql = "INSERT INTO roles (name) VALUES ($1)"
    updatePsql = "UPDATE roles SET name = $1, active = $2, updated_at = now() WHERE id = $3  RETURNING name, active, created_at, updated_at"
    updateMsql = "UPDATE roles SET name = $1, active = $2, updated_at = now() WHERE id = $3"
    delete     = "DELETE FROM roles WHERE id = $1"
    selectByID = "SELECT id, name, active, created_at, updated_at FROM roles WHERE id = $1"
    selectAll  = "SELECT id, name, active, created_at, updated_at FROM roles"
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
func (r *Rol) Create() {
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        createPostgresql(r)
    case conn.MYSQL:
        createMysql(r)
    }
}

// Update actualiza un registro en la BD
func (r *Rol) Update() {
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        updatePostgresql(r)
    case conn.MYSQL:
        updateMysql(r)
    }
}

// Delete borra un registro de la BD
func (r *Rol) Delete() {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(delete, r.ID)
    switch {
    case err == sql.ErrNoRows:
        log.Println("No rows deleted", r.ID)
    case err != nil:
        log.Fatal(err)
    }

    if i, _ := row.RowsAffected(); i == 1 {
        *r = Rol{}
    }
}

// GetByID Obtiene un rol por su ID
func (r *Rol) GetByID() {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(selectByID, r.ID).Scan(&r.ID, &r.Name, &r.Active, &r.CreatedAt, &r.UpdatedAt)
    switch {
    case err == sql.ErrNoRows:
        log.Println("No rows found with ID", r.ID)
    case err != nil:
        log.Fatal(err)
    }
}

// GetAll Obtiene todos los registros de la tabla
func (r *Rol) GetAll() []Rol {
    var roles = make([]Rol, 0)
    db := conn.Get()
    defer db.Close()

    rows, err := db.Query(selectAll)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var r Rol
        err := rows.Scan(&r.ID, &r.Name, &r.Active, &r.CreatedAt, &r.UpdatedAt)
        if err != nil {
            log.Fatal(err)
        }
        roles = append(roles, r)
    }
    return roles
}

func createPostgresql(r *Rol) {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(insertPsql, r.Name).Scan(&r.ID, &r.Name, &r.Active, &r.CreatedAt, &r.UpdatedAt)
    switch {
    case err == sql.ErrNoRows:
        log.Print("No rows inserted", r.Name)
    case err != nil:
        log.Fatal(err)
    }
}

func updatePostgresql(r *Rol) {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(updatePsql, r.Name, r.Active, r.ID).Scan(&r.Name, &r.Active, &r.CreatedAt, &r.UpdatedAt)
    switch {
    case err == sql.ErrNoRows:
        log.Println("No rows updated", r.ID, r.Name, r.Active)
    case err != nil:
        log.Fatal(err)
    }
}

func createMysql(r *Rol) {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(insertMsql, r.Name)
    if err != nil {
        log.Fatal(err)
    }
    id, err := row.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }
    r.ID = uint(id)
    r.GetByID()
}

func updateMysql(r *Rol) {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(updateMsql, r.Name, r.Active, r.ID)
    if err != nil {
        log.Fatal(err)
    }
    _, err = row.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    r.GetByID()
}
