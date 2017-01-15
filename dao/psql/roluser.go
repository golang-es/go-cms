package psql

import "github.com/golang-es/go-cms/models"

type RolUserDAOPSQL struct{}

// InsertRolUser inserta un regsitro en la bd
func (r RolUserDAOPSQL) InsertRolUser(ru *models.RoleUser) error {
	query := "INSERT INTO role_user (role_id, user_id) VALUES ($1, $1) RETURNING id, created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, ru.RoleID, ru.UserID).Scan(&ru.ID, &ru.CreatedAt, &ru.UpdatedAt)
	return err
}

// UpdateRolUser actualiza un registro en la bd
func (r RolUserDAOPSQL) UpdateRolUser(ru *models.RoleUser) error {
	query := "UPDATE role_user SET role_id = $1, user_id = $2, updated_at = now() WHERE id = $3 RETURNING created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, ru.RoleID, ru.UserID, ru.ID).Scan(&ru.CreatedAt, &ru.UpdatedAt)
	return err
}

// DeleteRolUser borra un registro en la bd
func (r RolUserDAOPSQL) DeleteRolUser(ru *models.RoleUser) error {
	query := "DELETE FROM role_user WHERE id = $1"
	db := get()
	defer db.Close()

	row, err := db.Exec(query, ru.ID)
	if err != nil {
		return err
	}

	if i, _ := row.RowsAffected(); i == 1 {
		ru = &models.RoleUser{}
	}
	return nil
}

// GetByIDRolUser obtiene un registro de la bd por el id
func (r RolUserDAOPSQL) GetByIDRolUser(id int) (*models.RoleUser, error) {
	query := "SELECT id, role_id, user_id, created_at, updated_at FROM role_user WHERE id = $1"
	ru := &models.RoleUser{}
	db := get()
	defer db.Close()

	err := db.QueryRow(query, id).Scan(&ru.ID, &ru.RoleID, &ru.UserID, &ru.CreatedAt, &ru.UpdatedAt)
	return ru, err
}

// GetAllRolUser obtiene todos los registros de la bd
func (r RolUserDAOPSQL) GetAllRolUser() ([]models.RoleUser, error) {
	query := "SELECT id, role_id, user_id, created_at, updated_at FROM role_user"
	var roleByUsers = make([]models.RoleUser, 0)
	db := get()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var roleByUser models.RoleUser
		err := rows.Scan(&roleByUser.ID, &roleByUser.RoleID, &roleByUser.UserID, &roleByUser.CreatedAt, &roleByUser.UpdatedAt)
		if err != nil {
			return roleByUsers, err
		}
		roleByUsers = append(roleByUsers, roleByUser)
	}
	return roleByUsers, nil
}
