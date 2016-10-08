package routers

import (
	"github.com/gorilla/mux"
	"github.com/golang-es/go-cms/controllers"
)

func SetPostRouters(router *mux.Router) *mux.Router  {
	//Ruta Publica
	router.HandleFunc("/post/{slug}", controllers.ShowPost).Methods("GET")
	//Rutas que necesitan autentificación
	router.HandleFunc("/admin/post/new", controllers.NewPost).Methods("GET")
	router.HandleFunc("/admin/post/edit/{id}", controllers.EditPost).Methods("GET")
	router.HandleFunc("/admin/post/save", controllers.SaveNewPost).Methods("POST")
	router.HandleFunc("/admin/post/save/{id}", controllers.SaveEditedPost).Methods("PUT")
	router.HandleFunc("/admin/post/delete/{id}", controllers.DeletePost).Methods("DELETE")

	return router

}
