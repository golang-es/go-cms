package psql

import "github.com/golang-es/go-cms/models"

type RolDAOPSQL struct{}

// InsertRol Inserta un rol en la BD
func (r RolDAOPSQL) InsertRol(rol *models.Rol) error {
	query := "INSERT INTO roles (name) VALUES ($1) RETURNING id, active, created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, rol.Name).Scan(&rol.ID, &rol.Active, &rol.CreatedAt, &rol.UpdatedAt)
	return err
}

// UpdateRol Actualiza un rol en la BD
func (r RolDAOPSQL) UpdateRol(rol *models.Rol) error {
	query := "UPDATE roles SET name = $1, active = $2, updated_at = now() WHERE id = $3  RETURNING created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, rol.Name, rol.Active, rol.ID).Scan(&rol.CreatedAt, &rol.UpdatedAt)
	return err
}

// DeleteRol Borra un rol en la BD
func (r RolDAOPSQL) DeleteRol(rol *models.Rol) error {
	query := "DELETE FROM roles WHERE id = $1"
	db := get()
	defer db.Close()

	row, err := db.Exec(query, rol.ID)
	if err != nil {
		return err
	}

	if i, _ := row.RowsAffected(); i == 1 {
		rol = &models.Rol{}
	}
	return nil
}

// GetByIDRol Obtiene un rol de la BD
func (r RolDAOPSQL) GetByIDRol(id uint) (*models.Rol, error) {
	query := "SELECT id, name, active, created_at, updated_at FROM roles WHERE id = $1"
	rol := &models.Rol{}
	db := get()
	defer db.Close()

	err := db.QueryRow(query, id).Scan(&rol.ID, &rol.Name, &rol.Active, &rol.CreatedAt, &rol.UpdatedAt)
	return rol, err
}

// GetAllRol Obtiene todos los roles de la BD
func (r RolDAOPSQL) GetAllRol() ([]models.Rol, error) {
	query := "SELECT id, name, active, created_at, updated_at FROM roles"
	var roles = make([]models.Rol, 0)
	db := get()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rol models.Rol
		err := rows.Scan(&rol.ID, &rol.Name, &rol.Active, &rol.CreatedAt, &rol.UpdatedAt)
		if err != nil {
			return roles, err
		}
		roles = append(roles, rol)
	}
	return roles, nil
}
