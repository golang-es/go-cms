package executedao

import (
	"github.com/golang-es/go-cms/dao/configuration"
	"github.com/golang-es/go-cms/dao/interfaces"
	"github.com/golang-es/go-cms/dao/psql"
	"log"
	"sync"
)

var (
	once   sync.Once
	metadataDAO interfaces.MetadataDAO
	rolDAO interfaces.RolDAO
)

func init() {
	once.Do(func() {
		initDAO()
	})
}

// initDAO inicia los DAO dependiendo de la configuración de la conexión
func initDAO() {
	log.Println("Se inician los DAO...")
	switch configuration.Config.Engine {
	case "postgresql":
		rolDAO = psql.RolDAOPSQL{}
		metadataDAO = psql.MetadataDAOPSQL{}
	default:
		log.Fatal("No existe el motor de persistencia solicitado")
	}
}
