package connection

import (
    "encoding/json"
    "log"
    "os"
)

type configuration struct {
    Engine   string
    Server   string
    Port     string
    User     string
    Password string
    Database string
    Sslmode  string
}

// Lee la información de un archivo de configuración json config.json
func getConfigurationFile() {
    file, err := os.Open("../config.json")
    if err != nil {
        log.Fatal(err)
    }
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    if err != nil {
        log.Fatal(err)
    }
    setEngine(config.Engine)
}

func setConfigurationEngine() {
    s.setServer(config.Server)
    s.setPort(config.Port)
    s.setUser(config.User)
    s.setPassword(config.Password)
    s.setDataBase(config.Database)
    s.setSslMode(config.Sslmode)
}
