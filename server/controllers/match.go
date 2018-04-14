package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"net/http"
	"time"
)

// Ensure MatchController implements IController.
var _ core.IController = &MatchController{}

// MatchController is the structure for the controller of matches
type MatchController struct {
	core.Controller
	MatchService models.IMatchService
}

// LoadRoutes loads all the routes of matches
func (controller *MatchController) LoadRoutes(r *gin.RouterGroup) {
	router := r.Group("/matches")

	router.POST("/:userID/shoot", controller.Shoot)
}

// Shoot is a function to register a user shoot
func (controller *MatchController) Shoot(c *gin.Context) {
	// Check if own user
	userID, err := controller.GetRequestID(c, "userID")
	if err != nil {
		controller.HandleError(c, err)
		return
	}
	sessionID, isLogged := c.Get("userID")
	if !isLogged {
		controller.HandleError(c, core.ErrorNotLogged)
		return
	}
	if userID != sessionID {
		controller.HandleError(c, core.ErrorNoPermission)
		return
	}

	// Get user match
	match, err := controller.MatchService.GetActiveMatchByUserID(userID, core.GetConfiguration().MatchTimeout)
	if match == nil {
		controller.HandleError(c, core.ErrorNotInMatch)
		return
	}

	// Set shoot time
	if match.UserOneID == userID {
		match.UserOneShootTime = time.Now()
	} else if match.UserTwoID == userID {
		match.UserTwoShootTime = time.Now()
	} else {
		// How are we here?
		controller.HandleError(c, core.ErrorInternalServerError)
		return
	}

	// Save the match
	err = controller.MatchService.Update(match)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
