package routers

import (
	"github.com/golang-es/go-cms/common"
	"github.com/golang-es/go-cms/controllers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAuthMiddleware(router *mux.Router) *mux.Router {
	//Rutas que necesitan autentificación
	//Agregamos los Middleware
	router.PathPrefix("/admin").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(router),
	))
	return router

}
