package main

import (
	"fmt"

	"github.com/golang-es/go-cms/dao/executedao"
)

func main() {
	rol, _ := executedao.GetByIDRol(1)
	fmt.Println(rol)
}
