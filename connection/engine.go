package connection

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
