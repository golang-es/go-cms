package psql

import "github.com/golang-es/go-cms/models"

type ModuleDAOPSQL struct{}

// InsertModule inserta un registro en la BD
func (m ModuleDAOPSQL) InsertModule(module *models.Module) error {
	query := "INSERT INTO modules (name, description) VALUES ($1, $2) RETURNING id, created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, module.Name, module.Description).Scan(&module.ID, &module.CreatedAt, &module.UpdatedAt)
	return err
}

// UpdateModule actualiza un registro en la BD
func (m ModuleDAOPSQL) UpdateModule(module *models.Module) error {
	query := "UPDATE modules SET name = $1, description = $2, updated_at = now() WHERE id = $3  RETURNING created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, module.Name, module.Description, module.ID).Scan(&module.CreatedAt, &module.UpdatedAt)
	return err
}

// DeleteModule borra un registro en la BD
func (m ModuleDAOPSQL) DeleteModule(module *models.Module) error {
	query := "DELETE FROM modules WHERE id = $1"
	db := get()
	defer db.Close()

	row, err := db.Exec(query, module.ID)
	if err != nil {
		return err
	}

	if i, _ := row.RowsAffected(); i == 1 {
		module = &models.Module{}
	}
	return nil
}

// GetByIDModule consulta un registro por id de la BD
func (m ModuleDAOPSQL) GetByIDModule(id int) (*models.Module, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM modules WHERE id = $1"
	model := &models.Module{}
	db := get()
	defer db.Close()

	err := db.QueryRow(query, id).Scan(&model.ID, &model.Name, &model.Description, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}

// GetAllModule consulta todos los registros de la BD
func (m ModuleDAOPSQL) GetAllModule() ([]models.Module, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles"
	var modules = make([]models.Module, 0)
	db := get()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var module models.Module
		err := rows.Scan(&module.ID, &module.Name, &module.Description, &module.CreatedAt, &module.UpdatedAt)
		if err != nil {
			return modules, err
		}
		modules = append(modules, module)
	}
	return modules, nil
}
