package executedao

import (
	"github.com/golang-es/go-cms/dao/configuration"
	"github.com/golang-es/go-cms/dao/interfaces"
	"github.com/golang-es/go-cms/dao/psql"
	"log"
	"sync"
)

var (
	once            sync.Once
	metadataDAO     interfaces.MetadataDAO
	metadataTypeDAO interfaces.MetadataTypeDAO
	moduleDAO       interfaces.ModuleDAO
	moduleRolDAO    interfaces.ModuleRolDAO
	postDAO         interfaces.PostDAO
	rolDAO          interfaces.RolDAO
	rolUserDAO		interfaces.RolUserDAO
	userDAO			interfaces.UserDAO
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
		metadataDAO = psql.MetadataDAOPSQL{}
		metadataTypeDAO = psql.MetadataTypeDAOPSQL{}
		moduleDAO = psql.ModuleDAOPSQL{}
		moduleRolDAO = psql.ModuleRolDAOPSQL{}
		postDAO = psql.PostDAOPSQL{}
		rolDAO = psql.RolDAOPSQL{}
		rolUserDAO = psql.RolUserDAOPSQL{}
		userDAO = psql.UserDAOPSQL{}

	default:
		log.Fatal("No existe el motor de persistencia solicitado")
	}
}
