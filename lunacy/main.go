package main

import (
	"fmt"

	_ "github.com/golang-es/go-cms/connection"
	"github.com/golang-es/go-cms/models"
)

func main() {
	rol := models.Rol{ID: 3}
	rol.GetByID()
	fmt.Println(rol)
}
