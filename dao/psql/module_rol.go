package psql

import "github.com/golang-es/go-cms/models"

type ModuleRolDAOPSQL struct{}

// InsertModuleRol inserta un registro en la BD
func (m ModuleRolDAOPSQL) InsertModuleRol(mr *models.ModuleRol) error {
	query := "INSERT INTO module_rol (module_id, role_id, append, modify, remove, read) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, mr.ModuleID, mr.RoleID, mr.Append, mr.Modify, mr.Remove, mr.Read).Scan(&mr.ID, &mr.CreatedAt, &mr.UpdatedAt)
	return err
}

// UpdateModuleRol actualiza un registro en la BD
func (m ModuleRolDAOPSQL) UpdateModuleRol(mr *models.ModuleRol) error {
	query := "UPDATE module_rol SET append = $1, modify = $2, remove = $3, read = $4, updated_at = now() WHERE id = $5 RETURNING module_id, role_id, created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, mr.Append, mr.Modify, mr.Remove, mr.Read, mr.ID).Scan(&mr.CreatedAt, &mr.UpdatedAt)
	return err
}

// DeleteModuleRol borra un registro en la BD
func (m ModuleRolDAOPSQL) DeleteModuleRol(mr *models.ModuleRol) error {
	query := "DELETE FROM module_rol WHERE id = $1"
	db := get()
	defer db.Close()

	row, err := db.Exec(query, mr.ID)
	if err != nil {
		return err
	}

	if i, _ := row.RowsAffected(); i == 1 {
		mr = &models.ModuleRol{}
	}
	return nil
}

// GetByIDModuleRol obtiene un registro en la BD por id
func (m ModuleRolDAOPSQL) GetByIDModuleRol(id int) (*models.ModuleRol, error) {
	query := "SELECT id, module_id, role_id, append, modify, remove, read, created_at, updated_at FROM module_rol WHERE id = $1"
	mr := &models.ModuleRol{}
	db := get()
	defer db.Close()

	err := db.QueryRow(query, id).Scan(&mr.ModuleID, &mr.RoleID, &mr.Append, &mr.Modify, &mr.Remove, &mr.Read, &mr.CreatedAt, &mr.UpdatedAt)
	return mr, err
}

// GetAllModuleRol selecciona todos los registros de la BD
func (m ModuleRolDAOPSQL) GetAllModuleRol() ([]models.ModuleRol, error) {
	query := "SELECT id, module_id, role_id, append, modify, remove, read, created_at, updated_at FROM module_rol"
	var modulesByRoles = make([]models.ModuleRol, 0)
	db := get()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var moduleByRol models.ModuleRol
		err := rows.Scan(&moduleByRol.ID, &moduleByRol.ModuleID, &moduleByRol.RoleID, &moduleByRol.Append, &moduleByRol.Modify, &moduleByRol.Remove, &moduleByRol.Read, &moduleByRol.CreatedAt, &moduleByRol.UpdatedAt)
		if err != nil {
			return modulesByRoles, err
		}
		modulesByRoles = append(modulesByRoles, moduleByRol)
	}
	return modulesByRoles, nil
}
