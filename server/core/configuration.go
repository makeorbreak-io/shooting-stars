package core

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

var config *Configuration

// Configuration is the struct that holds the application configuration
type Configuration struct {
	AppName               string `json:"app_name"`
	BaseURL               string `json:"base_url"`
	Port                  int    `json:"port"`
	Production            bool   `json:"production"`
	DatabaseType          string `json:"database_type"`
	DatabaseConnection    string `json:"database_connection"`
	MaxLocationLastUpdate uint   `json:"maxLocationLastUpdate"`
	MatchTimeout          uint   `json:"matchTimeout"`
	MatchMakingInterval   uint   `json:"matchMakingInterval"`
	LeaderBoardSize       uint   `json:"leaderBoardSize"`
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
		log.Printf("Could not load the configuration. Error: %v", err)
		return nil
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(config)
	if err != nil {
		log.Printf("Could not decode the configuration. Error: %v", err)
		return nil
	}

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
	envProduction := os.Getenv("PRODUCTION")
	if envProduction != "" {
		production, _ := strconv.ParseBool(envProduction)
		config.Production = production
	}
	envDatabase := os.Getenv("DATABASE_URL")
	if envDatabase != "" {
		config.DatabaseConnection = envDatabase
	}

	return config
}
