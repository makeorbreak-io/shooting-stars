package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"net/http"
)

// Ensure AuthController implements IController.
var _ core.IController = &AuthController{}

// AuthController is the structure for the controller of authentication
type AuthController struct {
	core.Controller
	AuthService models.IAuthService
	UserService models.IUserService
}

// LoadRoutes loads all the routes of authentication
func (controller *AuthController) LoadRoutes(r *gin.RouterGroup) {
	router := r.Group("/auth")

	router.POST("/login", controller.Login)
	router.GET("/register", controller.Register)
}

// Login is a function that validates a login and returns an authentication token
func (controller *AuthController) Login(c *gin.Context) {
	// Get request information
	var request models.LoginRequest
	err := c.ShouldBind(&request)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	// Verify request
	if request.Email == "" ||
		request.Password == "" {
		controller.HandleError(c, core.ErrorBadRequest)
		return
	}

	// Check if user exists
	user, err := controller.UserService.GetByEmail(request.Email)
	if err != nil {
		controller.HandleError(c, core.ErrorNotFound)
		return
	}

	// Verify the password
	err = controller.AuthService.ValidateLogin(request.Password, user)
	if err != nil {
		controller.HandleError(c, core.ErrorBadLogin)
		return
	}

	// Generate authentication token
	authToken := controller.AuthService.GenerateAuthToken(user.ID)

	c.JSON(http.StatusOK, authToken)
}

// Register is a function that creates a user registration
func (controller *AuthController) Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	_, err = controller.UserService.Create(&user)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}