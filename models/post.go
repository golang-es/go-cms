package models

import (
	"errors"
	"time"

	conn "github.com/golang-es/go-cms/connection"
)

const (
	postInsertPsql = "INSERT INTO posts (user_id, title, slug, content, published_at, poster, banner) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at"
	postUpdatePsql = "UPDATE posts SET user_id = $1, title = $2, slug = $3, content = $4, published_at = $5, poster = $6, banner = $7, updated_at = now() WHERE id = $8 RETURNING created_at, updated_at"
	postDelete     = "DELETE FROM post WHERE id = $1"
	postSelectByID = "SELECT id, user_id, title, slug, content, published_at, poster, banner, created_at, updated_at FROM posts WHERE id = $1"
	postSelectAll  = "SELECT id, user_id, title, slug, content, published_at, poster, banner, created_at, updated_at FROM posts"
)

// Post post que publican los usuarios del cms
type Post struct {
	ID          uint
	UserID      uint
	Title       string
	Slug        string
	Content     string
	PublishedAt time.Time
	Poster      string
	Banner      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Create crea un registro en la BD
func (p *Post) Create() error {
	var err error
	switch conn.GetNameEngine() {
	case conn.POSTGRESQL:
		err = createPostPostgresql(p)
	case conn.MYSQL:
		err = errors.New("Not supported yet.")
	}
	return err
}

// Update actualiza un registro en la BD
func (p *Post) Update() error {
	var err error
	switch conn.GetNameEngine() {
	case conn.POSTGRESQL:
		err = updatePostPostgresql(p)
	case conn.MYSQL:
		err = errors.New("Not supported yet.")
	}
	return err
}

// Delete Borra un registro de la BD
func (p *Post) Delete() error {
	db := conn.Get()
	defer db.Close()

	row, err := db.Exec(postDelete, p.ID)
	if err != nil {
		return err
	}

	if i, _ := row.RowsAffected(); i == 1 {
		*p = Post{}
	}
	return nil
}

// GetByID obtiene un registro de la BD
func (p *Post) GetByID() error {
	db := conn.Get()
	defer db.Close()

	err := db.QueryRow(postSelectByID, p.ID).Scan(&p.ID, &p.UserID, &p.Title, &p.Slug, &p.Content, &p.PublishedAt, &p.Poster, &p.Banner, &p.CreatedAt, &p.UpdatedAt)
	return err
}

// GetAll obtiene todos los registros de la BD
func (p *Post) GetAll() ([]Post, error) {
	var posts []Post
	db := conn.Get()
	defer db.Close()

	rows, err := db.Query(postSelectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Slug, &post.Content, &post.PublishedAt, &post.Poster, &post.Banner, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Esta función inserta un nuevo registro en POSTGRESQL
func createPostPostgresql(p *Post) error {
	db := conn.Get()
	defer db.Close()

	err := db.QueryRow(postInsertPsql, &p.UserID, &p.Title, &p.Slug, &p.Content, &p.PublishedAt, &p.Poster, &p.Banner).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	return err
}

// Esta función actualiza un registro en POSTGRESQL
func updatePostPostgresql(p *Post) error {
	db := conn.Get()
	defer db.Close()

	err := db.QueryRow(postUpdatePsql, &p.UserID, &p.Title, &p.Slug, &p.Content, &p.PublishedAt, &p.Poster, &p.Banner, &p.ID).Scan(&p.CreatedAt, &p.UpdatedAt)
	return err
}
