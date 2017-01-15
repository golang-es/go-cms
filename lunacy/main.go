package main

import (
	"fmt"

	"github.com/golang-es/go-cms/dao/executedao"
	"github.com/golang-es/go-cms/models"
)

func main() {
	newRol := &models.Rol{Name: "Desde DAO"}
	executedao.InsertRol(newRol)
	fmt.Println(newRol)
	rol, _ := executedao.GetByIDRol(1)
	rols, _ := executedao.GetAllRol()
	fmt.Println(rol)
	fmt.Println(rols)
}
