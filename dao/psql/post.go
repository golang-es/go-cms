package psql

import "github.com/golang-es/go-cms/models"

type PostDAOPSQL struct{}

// InsertPost inserta un registro en la BD
func (p PostDAOPSQL) InsertPost(post *models.Post) error {
	query := "INSERT INTO posts (user_id, title, slug, content, published_at, poster, banner) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, &post.UserID, &post.Title, &post.Slug, &post.Content, &post.PublishedAt, &post.Poster, &post.Banner).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	return err
}

// UpdatePost actualiza un registro en la BD
func (p PostDAOPSQL) UpdatePost(post *models.Post) error {
	query := "UPDATE posts SET user_id = $1, title = $2, slug = $3, content = $4, published_at = $5, poster = $6, banner = $7, updated_at = now() WHERE id = $8 RETURNING created_at, updated_at"
	db := get()
	defer db.Close()

	err := db.QueryRow(query, &post.UserID, &post.Title, &post.Slug, &post.Content, &post.PublishedAt, &post.Poster, &post.Banner, &post.ID).Scan(&post.CreatedAt, &post.UpdatedAt)
	return err
}

// DeletePost borra un registro en la BD
func (p PostDAOPSQL) DeletePost(post *models.Post) error {
	query := "DELETE FROM post WHERE id = $1"
	db := get()
	defer db.Close()

	row, err := db.Exec(query, post.ID)
	if err != nil {
		return err
	}

	if i, _ := row.RowsAffected(); i == 1 {
		post = &models.Post{}
	}
	return nil
}

// GetByIDPost obtiene un registro por id de la BD
func (p PostDAOPSQL) GetByIDPost(id int) (*models.Post, error) {
	query := "SELECT id, user_id, title, slug, content, published_at, poster, banner, created_at, updated_at FROM posts WHERE id = $1"
	post := &models.Post{}
	db := get()
	defer db.Close()

	err := db.QueryRow(query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Slug, &post.Content, &post.PublishedAt, &post.Poster, &post.Banner, &post.CreatedAt, &post.UpdatedAt)
	return post, err
}

// GetAllPost obtiene todos los registros de la BD
func (p PostDAOPSQL) GetAllPost() ([]models.Post, error) {
	query := "SELECT id, user_id, title, slug, content, published_at, poster, banner, created_at, updated_at FROM posts"
	var posts []models.Post
	db := get()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Slug, &post.Content, &post.PublishedAt, &post.Poster, &post.Banner, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
