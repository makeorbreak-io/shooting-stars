package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"golang.org/x/net/websocket"
	"log"
)

// Ensure MatchController implements IController.
var _ core.IController = &WebSocketController{}

// WebSocketController is the structure for the controller of web socket
type WebSocketController struct {
	core.Controller
	AuthService models.IAuthService
}

// LoadRoutes loads all the routes of matches
func (controller *WebSocketController) LoadRoutes(r *gin.RouterGroup) {
	router := r.Group("/websocket")

	router.GET("/:token", controller.Join)
}

// Join is a function so that a user joins a match
func (controller *WebSocketController) Join(c *gin.Context) {
	// Authenticate request
	token, err := controller.GetRequestParam(c, "token")
	if token == "" || err != nil {
		controller.HandleError(c, core.ErrorNotLogged)
		return
	}

	_, err = controller.AuthService.GetUserIDByToken(token)
	if err != nil {
		controller.HandleError(c, core.ErrorNotLogged)
		return
	}

	// Handle web socket
	handler := websocket.Handler(controller.WebSocketHandler)
	handler.ServeHTTP(c.Writer, c.Request)
}

// WebSocketHandler is the handler for web sockets
func (controller *WebSocketController) WebSocketHandler(ws *websocket.Conn) {
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
