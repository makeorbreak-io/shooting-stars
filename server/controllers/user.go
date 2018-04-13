package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"net/http"
)

// Ensure UserController implements IController.
var _ core.IController = &UserController{}

// UserController is the structure for the controller of users
type UserController struct {
	core.Controller
	UserService models.IUserService
}

// LoadRoutes loads all the routes of users
func (controller *UserController) LoadRoutes(r *gin.RouterGroup) {
	router := r.Group("/users")

	// Educations
	router.GET("", controller.GetAll)
	router.GET("/:id", controller.Get)
}

// GetAll is a method to get all users
func (controller *UserController) GetAll(c *gin.Context) {
	// Get all educations
	educations, err := controller.UserService.GetAll()
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, educations)
}

// Get is a method to get an user
func (controller *UserController) Get(c *gin.Context) {
	id, err := controller.GetRequestID(c)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	education, err := controller.UserService.Get(id)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, education)
}

// Create is a method to create a user upon a request
func (controller *UserController) Create(c *gin.Context) {
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

// Edit is a method to edit a education upon a request
func (controller *UserController) Edit(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	// Check if editing the education of the request
	id, err := controller.GetRequestID(c)
	if err != nil {
		controller.HandleError(c, err)
		return
	}
	if id != user.ID {
		controller.HandleError(c, core.ErrorBadRequest)
	}

	// Update the education
	err = controller.UserService.Update(&user)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete is a method to delete a education upon a request
func (controller *UserController) Delete(c *gin.Context) {
	id, err := controller.GetRequestID(c)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	err = controller.UserService.Delete(id)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
