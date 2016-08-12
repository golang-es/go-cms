// Package models contiene el modelo de la App
package models

// Model es la interface para acceder a los modelos
type Model interface {
    Create() error
    Update() error
    Delete() error
    GetByID() error
    GetAll() ([]Model, error)
}
