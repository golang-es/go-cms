package main

import (
	"fmt"

	"github.com/golang-es/go-cms/models"
)

func main() {
	rol := models.Rol{ID: 4}
	rol.GetByID()
	fmt.Println(rol)
}
