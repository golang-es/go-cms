package connection

import "database/sql"

// Source Interface que permite obtener una estructura que conecte
// a una fuente de datos.
type source interface {
    connection() *sql.DB
    setServer(server string)
    setPort(port string)
    setUser(user string)
    setPassword(password string)
    setDataBase(database string)
    setSslMode(sslmode string)
}
