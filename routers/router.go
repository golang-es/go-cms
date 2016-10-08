package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router  {
	router := mux.NewRouter().StrictSlash(false)
	//Ruta para la  autentificación o registro de usuarios
	router = SetUserRouter(router)
	//Ruta para la administracón de los posts
	router = SetPostRouters(router)
	//Middleware de autentificación para la ruta /admin
	router = SetAuthMiddleware(router)
	//Retornamos el router configurado
	return router
}
