package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"golang.org/x/crypto/bcrypt"
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
	router.POST("/register", controller.Register)
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
	authToken, err := controller.AuthService.GenerateAuthToken(user.ID)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, authToken)
}

// Register is a function that creates a user registration
func (controller *AuthController) Register(c *gin.Context) {
	var request models.RegisterRequest
	err := c.ShouldBind(&request)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	// Verify request
	if request.Name == "" ||
		request.Email == "" ||
		request.Password == "" ||
		request.ConfirmPassword == "" ||
		request.Gender == "" {
		controller.HandleError(c, core.ErrorBadRequest)
		return
	}
	if request.Password != request.ConfirmPassword {
		controller.HandleError(c, core.ErrorPasswordsDontMatch)
		return
	}
	if request.Gender != "M" && request.Gender != "F" {
		controller.HandleError(c, core.ErrorInvalidGender)
		return
	}

	// Create user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		controller.HandleError(c, err)
		return
	}
	newUser := models.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
		BirthDate:    request.BirthDate,
		Gender:       request.Gender,
	}

	_, err = controller.UserService.Create(&newUser)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	// Generate authentication token
	authToken, err := controller.AuthService.GenerateAuthToken(newUser.ID)
	if err != nil {
		c.Status(http.StatusCreated)
		return
	}

	c.JSON(http.StatusCreated, authToken)
}
