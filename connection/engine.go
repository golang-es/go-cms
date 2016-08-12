package connection

import "log"

type engine struct {
    server   string
    port     string
    user     string
    password string
    database string
    sslmode  string
}

func (e *engine) setServer(server string) {
    e.server = server
}

func (e *engine) setPort(port string) {
    e.port = port
}

func (e *engine) setUser(user string) {
    e.user = user
}

func (e *engine) setPassword(password string) {
    e.password = password
}

func (e *engine) setDataBase(database string) {
    e.database = database
}

func (e *engine) setSslMode(sslmode string) {
    e.sslmode = sslmode
}

func setConfigurationEngine() {
    s.setServer(config.Server)
    s.setPort(config.Port)
    s.setUser(config.User)
    s.setPassword(config.Password)
    s.setDataBase(config.Database)
    s.setSslMode(config.Sslmode)
}

// Establece el motor al que se va a conectar
func setEngine(e string) {
    switch e {
    case POSTGRESQL:
        s = getInstancePostgresql()
    case MYSQL:
        s = getInstanceMysql()
    default:
        log.Fatal("No se puede registrar ese motor")
    }
    setConfigurationEngine()
}
