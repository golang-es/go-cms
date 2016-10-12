package main

import (
	"net/http"

	"github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1"
	"github.com/urfave/negroni"
)

func main() {
	r := render.New(render.Options{
		Directory:  "./",
		Extensions: []string{".html"},
		Layout:     "layout",
	})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "home", "Daniel")
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":1700")
}
