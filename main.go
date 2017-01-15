package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"

	"github.com/golang-es/go-cms/common"
	"github.com/golang-es/go-cms/database"
	"github.com/golang-es/go-cms/routers"
)

func main() {
	var seed string
	flag.StringVar(&seed, "seed", "no", "Inserta los registros iniciales "+
		"en la base de datos")
	flag.IntVar(&common.Port, "port", 1700, "Puerto para el servidor web")
	flag.Parse()

	if seed == "yes" {
		database.Seed()
		return
	}

	// Inicia las rutas
	router := routers.InitRoutes()

	// Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", common.Port),
		Handler: n,
	}

	log.Printf("Iniciado en http://localhost:%d", common.Port)
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución de go-cms")
}
