// Package connection Permite obtener una conexión a
// diferentes fuentes de datos
package connection

import (
    "database/sql"
    "sync"
)

// Variables que servirán de instancia para el patrón Singleton
var (
    config     configuration
    s          source
    p          *postgresql
    m          *mysql
    nameEngine string
    once       sync.Once
)

const (
    // POSTGRESQL nombre del motor de BD Postgresql
    POSTGRESQL = "postgresql"
    // MYSQL nombre del motor de BD Mysql
    MYSQL = "mysql"
)

func init() {
    config = configuration{}
    getConfigurationFile()
}

// Get devuelve la conexión a la BD
func Get() *sql.DB {
    return s.connection()
}

// GetNameEngine devuelve el nombre del motor de la BD
func GetNameEngine() string {
    return nameEngine
}

// Devuelve una única instancia de postgresql
func getInstancePostgresql() *postgresql {
    once.Do(func() {
        p = new(postgresql)
        nameEngine = "postgresql"
    })
    return p
}

// Devuelve una única instancia de mysql
func getInstanceMysql() *mysql {
    once.Do(func() {
        m = &mysql{}
        nameEngine = "mysql"
    })
    return m
}
