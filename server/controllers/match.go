package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

// Ensure MatchController implements IController.
var _ core.IController = &MatchController{}

// MatchController is the structure for the controller of matches
type MatchController struct {
	core.Controller
	UserService  models.IUserService
	MatchService models.IMatchService
}

// LoadRoutes loads all the routes of matches
func (controller *MatchController) LoadRoutes(r *gin.RouterGroup) {
	router := r.Group("/matches")

	router.POST("/:userID/shoot", controller.Shoot)
	router.POST("/:userID/join", controller.Join)
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
	currentTime := time.Now()
	if match.UserOneID == userID {
		match.UserOneShootTime = &currentTime
	} else if match.UserTwoID == userID {
		match.UserTwoShootTime = &currentTime
	} else {
		// How are we here?
		controller.HandleError(c, core.ErrorInternalServerError)
		return
	}

	// Check winner
	if match.WinnerID == nil {
		if match.UserOneShootTime == nil {
			match.WinnerID = &match.UserTwoID
		} else if match.UserTwoShootTime == nil {
			match.WinnerID = &match.UserOneID
		}

		err = controller.UserService.AddWin(*match.WinnerID)
		if err != nil {
			controller.HandleError(c, err)
			return
		}
	}

	// Save the match
	err = controller.MatchService.Update(match)
	if err != nil {
		controller.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, match)
}

// Join is a function so that a user joins a match
func (controller *MatchController) Join(c *gin.Context) {
	handler := websocket.Handler(controller.WebSocketHandler)
	handler.ServeHTTP(c.Writer, c.Request)
}

// WebSocketHandler is the handler for web sockets
func (controller *MatchController) WebSocketHandler(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			log.Println("Can't receive")
			break
		}

		log.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		log.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			log.Println("Can't send")
			break
		}
	}
}
