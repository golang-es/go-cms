package psql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alexyslozada/accounting-go/dao/configuration"
	// El paquete lib/pq es necesario para la conexión a postgresql
	_ "github.com/lib/pq"
)

// get Obtiene la conexión.
func get() *sql.DB {
	config := configuration.Config
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.User, config.Password, config.Server, config.Port, config.Database, config.Sslmode)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

