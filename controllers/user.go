package controllers

import (
	"net/http"
	"html/template"
	"github.com/golang-es/go-cms/models"
	"github.com/golang-es/go-cms/dao/executedao"
)

// Login procesa el formulario de login
func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Procesar el formulario de login
}

// Muestra la vista de registro
func ShowRegister(w http.ResponseWriter, r *http.Request) {
	renderUserRegister(w, nil)
}

// Registra procesa el formulario de registro
func Register(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	u.Name = r.FormValue("name")
	u.Lastname = r.FormValue("lastname")
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password")

	err := executedao.InsertUser(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderUserRegister(w, u)
}

func renderUserRegister(w http.ResponseWriter, i interface{}) {
	t, err := template.ParseFiles("./public/templates/user-register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return
	}

	t.Execute(w, i)
}
