package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"net/http"
)

// Ensure LocationController implements IController.
var _ core.IController = &LocationController{}

// LocationController is the structure for the controller of locations
type LocationController struct {
	core.Controller
	LocationService models.ILocationService
}

// LoadRoutes loads all the routes of authentication
func (controller *LocationController) LoadRoutes(r *gin.RouterGroup) {
	router := r.Group("/locations")

	router.POST("/:id", controller.Update)
}

// Update is a function that updates the location of a user
func (controller *LocationController) Update(c *gin.Context) {
	var location models.UpdateLocationRequest
	err := c.ShouldBindJSON(&location)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	// Check if editing own user
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

	// Update the location of the user
	previousLocation, err := controller.LocationService.Get(id)
	if previousLocation == nil ||
		err != nil {
		_, err = controller.LocationService.Create(id, location.Longitude, location.Longitude)
		if err != nil {
			controller.HandleError(c, err)
			return
		}
	} else {
		err = controller.LocationService.Update(id, location.Longitude, location.Longitude)
		if err != nil {
			controller.HandleError(c, err)
			return
		}
	}

	c.Status(http.StatusNoContent)
}
