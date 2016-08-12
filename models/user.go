package models

import (
    "errors"
    "time"

    conn "github.com/golang-es/go-cms/connection"
)

const (
    userInsertPsql = "INSERT INTO users (name, lastname, email, password) VALUES ($1, $2, $3, md5($4)) RETURNING id, created_at, updated_at"
    userUpdatePsql = "UPDATE users SET name = $1, lastname = $2, email = $3, password = md5($4), updated_at = now() WHERE id = $5 RETURNING created_at, updated_at"
    userDelete     = "DELETE FROM users WHERE id = $1"
    userSelectByID = "SELECT id, name, lastname, email, password, created_at, updated_at FROM users WHERE id = $1"
    userSelectAll  = "SELECT id, name, lastname, email, password, created_at, updated_at FROM users"
)

// User usuario del sistema
type User struct {
    ID        uint
    Name      string
    Lastname  string
    Email     string
    Password  string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Create crea un registro en la BD
func (u *User) Create() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = createUserPostgresql(u)
    case conn.MYSQL:
        err = errors.New("Not supported yet.")
    }
    return err
}

// Update actualiza un registro en la BD
func (u *User) Update() error {
    var err error
    switch conn.GetNameEngine() {
    case conn.POSTGRESQL:
        err = updateUserPostgresql(u)
    case conn.MYSQL:
        err = errors.New("Not supported yet.")
    }
    return err
}

// Delete borra un registro de la BD
func (u *User) Delete() error {
    db := conn.Get()
    defer db.Close()

    row, err := db.Exec(userDelete, u.ID)
    if err != nil {
        return err
    }

    if i, _ := row.RowsAffected(); i == 1 {
        *u = User{}
    }
    return nil
}

// GetByID Obtiene un registro por el ID
func (u *User) GetByID() error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(userSelectByID, u.ID).Scan(&u.ID, &u.Name, &u.Lastname, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
    return err
}

// GetAll Obtiene todos los registros
func (u *User) GetAll() ([]User, error) {
    var users = make([]User, 0)
    db := conn.Get()
    defer db.Close()

    rows, err := db.Query(userSelectAll)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
        if err != nil {
            return users, err
        }
        users = append(users, user)
    }
    return users, nil
}

func createUserPostgresql(u *User) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(userInsertPsql, u.Name, u.Lastname, u.Email, u.Password).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
    return err
}

func updateUserPostgresql(u *User) error {
    db := conn.Get()
    defer db.Close()

    err := db.QueryRow(userUpdatePsql, u.Name, u.Lastname, u.Email, u.Password, u.ID).Scan(&u.CreatedAt, &u.UpdatedAt)
    return err
}
