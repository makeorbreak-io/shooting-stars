package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"net/http"
)

// Ensure StatsController implements IController.
var _ core.IController = &StatsController{}

// StatsController is the structure for the controller of stats
type StatsController struct {
	core.Controller
	UserService models.IUserService
}

// LoadRoutes loads all the routes of stats
func (controller *StatsController) LoadRoutes(r *gin.RouterGroup) {
	router := r.Group("/stats")

	router.GET("/topWins", controller.GetUsersMostWins)
}

// GetUsersMostWins is a method to get the users with most wins
func (controller *StatsController) GetUsersMostWins(c *gin.Context) {
	// Check if logged in
	_, isLogged := c.Get("userID")
	if !isLogged {
		controller.HandleError(c, core.ErrorNotLogged)
		return
	}

	topUsers, err := controller.UserService.GetUsersMostWins(core.GetConfiguration().LeaderBoardSize)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	var users []*models.UserRank
	for idx, user := range topUsers {
		users = append(users, &models.UserRank{
			Name:   user.Name,
			Wins:   user.Wins,
			UserID: user.ID,
			Rank:   uint(idx + 1),
		})
	}

	c.JSON(http.StatusOK, users)
}
