package database

import (
	"log"

	"github.com/golang-es/go-cms/common"
	"github.com/golang-es/go-cms/dao/executedao"
	"github.com/golang-es/go-cms/models"
)

// Seed crea los registros iniciales de la base de datos
func Seed() {
	var err error

	role := models.Rol{Name: "admin"}
	log.Println("Creando el rol admin...")
	if err = executedao.InsertRol(&role); err != nil {
		log.Fatal(err)
	}
	log.Println("Rol admin creado")

	user := models.User{
		Name:     "Administrador",
		Lastname: "Gocms",
		Email:    "user@admin.dev",
		Password: common.PasswordSha256("admin"),
	}
	log.Println("Creando el usuario user@admin.dev...")
	if err = executedao.InsertUser(&user); err != nil {
		log.Fatal(err)
	}
	log.Println("Usuario user@admin.dev creado")

	roleUser := models.RoleUser{
		RoleID: role.ID,
		UserID: user.ID,
	}
	log.Println("Asignando el rol admin al usuario user@admin.dev...")
	if err = executedao.InsertRolUser(&roleUser); err != nil {
		log.Fatal(err)
	}
	log.Println("Rol admin asignado a user@admin.dev")
}
