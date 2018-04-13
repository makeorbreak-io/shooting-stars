package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // GORM needs this import in
	"github.com/makeorbreak-io/shooting-stars/server/controllers"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/services"
	"log"
	"os"
	"strconv"
)

func main() {
	// Flags
	var initDatabase = flag.Bool("initDB", false, `initialize the database.`)
	flag.Parse()

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

	// Create services
	authService := &services.AuthService{
		Database: database,
	}
	userService := &services.UserService{
		Database: database,
	}

	// Initialize database
	if *initDatabase {
		authService.CreateTable()
		userService.CreateTable()
	}

	// Load all controllers
	ctrls := []core.IController{
		&controllers.AuthController{
			AuthService: authService,
			UserService: userService,
		},
		&controllers.UserController{
			UserService: userService,
		},
	}
	routerBaseGroup := router.Group("")
	for _, controller := range ctrls {
		controller.LoadRoutes(routerBaseGroup)
	}

	// Start the server
	router.Run(":" + strconv.Itoa(config.Port))
}
