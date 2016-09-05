package psql

import "github.com/golang-es/go-cms/models"

type UserDAOPSQL struct {}

// InsertUser inserta un registro en la bd
func (u UserDAOPSQL) InsertUser(user *models.User) error {
	query := "INSERT INTO users (name, lastname, email, password) VALUES ($1, $2, $3, md5($4)) RETURNING id, created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, user.Name, user.Lastname, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	return err
}

// UpdateUser actualiza un registro en la bd
func (u UserDAOPSQL) UpdateUser(user *models.User) error {
	query := "UPDATE users SET name = $1, lastname = $2, email = $3, password = md5($4), updated_at = now() WHERE id = $5 RETURNING created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, user.Name, user.Lastname, user.Email, user.Password, user.ID).Scan(&user.CreatedAt, &user.UpdatedAt)
	return err
}

// DeleteUser borra un registro en la bd
func (u UserDAOPSQL) DeleteUser(user *models.User) error {
	query := "DELETE FROM users WHERE id = $1"
	db := get()
	defer db.Close()

	row, err := db.Exec(query, user.ID)
	if err != nil {
		return err
	}

	if i, _ := row.RowsAffected(); i == 1 {
		user = &models.User{}
	}
	return nil
}

// GetByIDUser obtiene un registro por el id en la bd
func (u UserDAOPSQL) GetByIDUser(id int) (*models.User, error) {
	query := "SELECT id, name, lastname, email, password, created_at, updated_at FROM users WHERE id = $1"
	user := &models.User{}
	db := get()
	defer db.Close()

	err := db.QueryRow(query, user.ID).Scan(&user.ID, &user.Name, &user.Lastname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

// GetAllUser obtiene todos los registros de la bd
func (u UserDAOPSQL) GetAllUser() ([]models.User, error) {
	query := "SELECT id, name, lastname, email, password, created_at, updated_at FROM users"
	var users = make([]models.User, 0)
	db := get()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
