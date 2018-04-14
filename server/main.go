package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // GORM needs this import in
	"github.com/makeorbreak-io/shooting-stars/server/controllers"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/middlewares"
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
	router.Use(middlewares.HandleExecutionErrors())
	router.Use(middlewares.HandleCors())
	// Router groups
	routerBaseGroup := router.Group("")
	routerPublicGroup := routerBaseGroup.Group("")
	routerAuthenticatedGroup := routerBaseGroup.Group("")
	routerAuthenticatedGroup.Use(middlewares.HandleAuthentication())

	// Setup database connection
	database, err := gorm.Open(config.DatabaseType, config.DatabaseConnection)
	if err != nil {
		log.Fatal("Could not connect to the database: " + err.Error())
		os.Exit(1)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "mob_" + defaultTableName
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
	locationService := &services.LocationService{
		Database: database,
	}
	matchService := &services.MatchService{
		Database: database,
	}
	userService := &services.UserService{
		Database: database,
	}

	// Initialize database
	if *initDatabase {
		authService.CreateTable()
		locationService.CreateTable()
		matchService.CreateTable()
		userService.CreateTable()
	}

	// Load all public controllers
	ctrls := []core.IController{
		&controllers.AuthController{
			AuthService: authService,
			UserService: userService,
		},
	}
	for _, controller := range ctrls {
		controller.LoadRoutes(routerPublicGroup)
	}

	// Load all authenticated controllers
	ctrls = []core.IController{
		&controllers.LocationController{
			LocationService: locationService,
		},
		&controllers.UserController{
			UserService: userService,
		},
	}
	for _, controller := range ctrls {
		controller.LoadRoutes(routerAuthenticatedGroup)
	}

	// Start the server
	router.Run(":" + strconv.Itoa(config.Port))
}
