package connection

import (
    "database/sql"
    "log"
    // El paquete es necesario para la conexión a mysql
    _ "github.com/go-sql-driver/mysql"
)

type mysql struct {
    engine
}

func (m *mysql) connection() *sql.DB {
    dataSourceName := m.user + ":" + m.password + "@tcp(" + m.server + ":" + m.port + ")/" + m.database
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }
    return db
}
