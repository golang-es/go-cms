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
	// La ruta debe ser sin ../
	file, err := os.Open("./config/config.json")
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
