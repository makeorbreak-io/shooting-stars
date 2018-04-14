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

	router.GET("/:id", controller.Get)
	router.POST("/:id", controller.Edit)
	router.DELETE("/:id", controller.Delete)
}

// Get is a method to get an user
func (controller *UserController) Get(c *gin.Context) {
	id, err := controller.GetRequestID(c)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	// Check if getting logged user
	sessionID, isLogged := c.Get("userID")
	if !isLogged {
		controller.HandleError(c, core.ErrorNotLogged)
		return
	}
	if id != sessionID {
		controller.HandleError(c, core.ErrorNoPermission)
		return
	}

	user, err := controller.UserService.Get(id)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Edit is a method to edit a education upon a request
func (controller *UserController) Edit(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	// Check if editing the user of the request
	id, err := controller.GetRequestID(c)
	if err != nil {
		controller.HandleError(c, err)
		return
	}
	sessionID, isLogged := c.Get("userID")
	if !isLogged {
		controller.HandleError(c, core.ErrorNotLogged)
		return
	}
	if id != sessionID {
		controller.HandleError(c, core.ErrorNoPermission)
		return
	}
	if id != user.ID {
		controller.HandleError(c, core.ErrorBadRequest)
		return
	}

	// Update the user
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

	// Check if deleting logged user
	sessionID, isLogged := c.Get("userID")
	if !isLogged {
		controller.HandleError(c, core.ErrorNotLogged)
		return
	}
	if id != sessionID {
		controller.HandleError(c, core.ErrorNoPermission)
		return
	}

	err = controller.UserService.Delete(id)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
