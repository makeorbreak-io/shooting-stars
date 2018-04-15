package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"github.com/makeorbreak-io/shooting-stars/server/tasks"
	"golang.org/x/net/websocket"
	"time"
)

// Ensure MatchController implements IController.
var _ core.IController = &WebSocketController{}

// WebSocketController is the structure for the controller of web socket
type WebSocketController struct {
	core.Controller
	AuthService     models.IAuthService
	MatchMakingTask *tasks.MatchMakingTask
}

// LoadRoutes loads all the routes of matches
func (controller *WebSocketController) LoadRoutes(r *gin.RouterGroup) {
	router := r.Group("/websocket")

	router.GET("/:token", controller.Join)
}

// Join is a function so that a user joins a match
func (controller *WebSocketController) Join(c *gin.Context) {
	// Handle web socket
	handler := websocket.Handler(controller.WebSocketHandler)
	handler.ServeHTTP(c.Writer, c.Request)
}

// WebSocketHandler is the handler for web sockets
func (controller *WebSocketController) WebSocketHandler(ws *websocket.Conn) {
	ws.SetReadDeadline(time.Now().Add(time.Second * 10))

	// Get authentication token
	var token string
	err := websocket.Message.Receive(ws, &token)
	if err != nil {
		websocket.Message.Send(ws, core.ErrorNotLogged.Error())
		ws.Close()
		return
	}

	// Authenticate request
	if token == "" {
		websocket.Message.Send(ws, core.ErrorNotLogged.Error())
		ws.Close()
		return
	}
	userID, err := controller.AuthService.GetUserIDByToken(token)
	if err != nil {
		websocket.Message.Send(ws, core.ErrorNotLogged.Error())
		ws.Close()
		return
	}

	websocket.Message.Send(ws, core.MessageOK)

	controller.MatchMakingTask.AddConnection(userID, ws)
}
