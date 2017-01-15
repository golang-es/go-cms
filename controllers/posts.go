package controllers

import (
	"fmt"
	"net/http"
)

// NewPost muestra el formulario para crear un post
func NewPost(w http.ResponseWriter, r *http.Request) {
	// TODO: Mostrar el formulario de nuevo post
	fmt.Println("Formulario para crear un nuevo post")
}

// ShowPost muestra un post
func ShowPost(w http.ResponseWriter, r *http.Request) {
	// TODO: Mostrar un post
	fmt.Println("Muestra un post")
}

// EditPost muestra el formulario para editar un post
func EditPost(w http.ResponseWriter, r *http.Request) {
	// TODO: Mostrar formulario para editar un post
	fmt.Println("Formulario para editar un post existente")
}

// SaveNewPost guarda los datos recibidos del formulario de crear post
func SaveNewPost(w http.ResponseWriter, r *http.Request) {
	// TODO: Procesar los datos del formulario de nuevo post
	fmt.Println("Guardar los datos de un nuevo post")
}

// SaveEditedPost guarda los datos recibidos del formulario de editar post
func SaveEditedPost(w http.ResponseWriter, r *http.Request) {
	// TODO: Guardar los datos editados de un post existente
	fmt.Println("Guardar los datos editados de un post existente")
}

// DeletePost elimina un post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// TODO: Eliminar un post existente
	fmt.Println("Elimina un post existente")
}
