package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // GORM needs this import in
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"log"
	"os"
	"strconv"
)

func main() {
	// Configuration
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not connect to the database: " + err.Error())
		os.Exit(1)
	}
	config := core.LoadConfiguration(currentDir + "/config.json")
	if config == nil {
		log.Fatal("Could not load the configuration file.")
		return
	}

	// Setup router
	router := gin.Default()

	// Setup database connection
	database, err := gorm.Open(config.DatabaseType, config.DatabaseConnection)
	if err != nil {
		log.Fatal("Could not connect to the database: " + err.Error())
		os.Exit(1)
	}

	// Verbose
	if config.Verbose {
		gin.SetMode(gin.DebugMode)
		database.LogMode(true)
	} else {
		gin.SetMode(gin.ReleaseMode)
		database.LogMode(false)
	}

	// Start the server
	router.Run(":" + strconv.Itoa(config.Port))
}
