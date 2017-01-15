package configuration

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var (
	// Config Configuración de la conexión
	Config         *configuration
	pathFileConfig = "./config/connection.json"
	once           sync.Once
)

func SetPathFileConfig(path string) {
	pathFileConfig = path
}

func init() {
	once.Do(func() {
		getConfigurationFile()
	})
}

// configuration Guarda la información del archivo de configuration.json
type configuration struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
	Sslmode  string
}

// Lee la información del archivo de configuración config.json
func getConfigurationFile() {
	log.Println("Se ha llamado getconfigurationfile...")
	// La ruta del archivo debe ser una carpeta
	// al mismo nivel del ejecutable principal (main)
	file, err := os.Open(pathFileConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal(err)
		return
	}
}
