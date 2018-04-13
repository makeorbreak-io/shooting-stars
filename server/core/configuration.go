package core

import (
	"encoding/json"
	"os"
	"strconv"
)

var config *Configuration

// Configuration is the struct that holds the application configuration
type Configuration struct {
	AppName            string `json:"app_name"`
	BaseURL            string `json:"base_url"`
	Port               int    `json:"port"`
	Verbose            bool   `json:"verbose"`
	DatabaseType       string `json:"database_type"`
	DatabaseConnection string `json:"database_connection"`
}

// GetConfiguration returns the configuration
func GetConfiguration() *Configuration {
	return config
}

// LoadConfiguration loads the configuration from a file
func LoadConfiguration(file string) *Configuration {
	config = new(Configuration)
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return nil
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(config)

	// Environment variables
	envBaseURL := os.Getenv("BASE_URL")
	if envBaseURL != "" {
		config.BaseURL = envBaseURL
	}
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port, _ := strconv.ParseInt(envPort, 10, 64)
		config.Port = int(port)
	}
	envDatabase := os.Getenv("DATABASE_URL")
	if envDatabase != "" {
		config.DatabaseConnection = envDatabase
	}

	return config
}
