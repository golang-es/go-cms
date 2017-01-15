package routers

import (
	"github.com/golang-es/go-cms/controllers"
	"github.com/gorilla/mux"
)

func SetPostRouters(router *mux.Router) *mux.Router {
	//Rutas que necesitan autentificación
	router.HandleFunc("/admin/posts/new", controllers.NewPost).Methods("GET")
	router.HandleFunc("/admin/posts/edit/{id}", controllers.EditPost).Methods("GET")
	router.HandleFunc("/admin/posts/save", controllers.SaveNewPost).Methods("POST")
	router.HandleFunc("/admin/posts/save/{id}", controllers.SaveEditedPost).Methods("PUT")
	router.HandleFunc("/admin/posts/delete/{id}", controllers.DeletePost).Methods("DELETE")
	//Ruta Publica
	router.HandleFunc("/{slug}", controllers.ShowPost).Methods("GET")

	return router

}
