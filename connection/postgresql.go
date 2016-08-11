package connection

import (
    "database/sql"
    "log"
    // El paquete lib/pq es necesario para la conexión a postgresql
    _ "github.com/lib/pq"
)

type postgresql struct {
    engine
}

func (p postgresql) connection() *sql.DB {
    dataSourceName := "postgres://" + p.user + ":" + p.password + "@" + p.server + ":" + p.port + "/" + p.database + "?sslmode=" + p.sslmode
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }
    return db
}
