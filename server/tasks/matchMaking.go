package tasks

import (
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

// Ensure MatchMakingTask implements ITask.
var _ ITask = &MatchMakingTask{}

// MatchMakingTask is the task related to the match making
type MatchMakingTask struct {
	LocationService models.ILocationService
	MatchService    models.IMatchService
	ticker          *time.Ticker
	quit            chan struct{}
}

// Web socket connections mapped from user to web socket
var connections map[uint]*websocket.Conn

// AddConnection adds a web socket connection associated to a given user
func (task *MatchMakingTask) AddConnection(userID uint, ws *websocket.Conn) {
	connections[userID] = ws
	log.Printf("User after adding connection %d", userID)
}

// Start is a function to start the task with a given interval between runs
func (task *MatchMakingTask) Start(interval uint) {
	task.ticker = time.NewTicker(time.Second * time.Duration(interval))
	task.quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-task.ticker.C:
				task.Run()
			case <-task.quit:
				task.ticker.Stop()
				return
			}
		}
	}()
}

// Stop stops the execution of the task
func (task *MatchMakingTask) Stop() {
	close(task.quit)
}

// Run executes the match making task
func (task *MatchMakingTask) Run() {
	activeUsersIDs, err := task.LocationService.GetActiveUsers(core.GetConfiguration().MaxLocationLastUpdate)
	if err != nil {
		log.Printf("Error when getting active users: %v", err)
		return
	}

	if len(activeUsersIDs) == 0 {
		log.Printf("Error no active users")
		return
	}

	for _, userID := range activeUsersIDs {
		nearestUsersLocations, err := task.LocationService.GetNearestActiveUserLocation(userID, core.GetConfiguration().MaxLocationLastUpdate)
		if err != nil {
			log.Printf("Error when getting nearest active users locations: %v", err)
			continue
		}

		// TODO: Remove me
		if ws, exists := connections[userID]; exists {
			log.Printf("Sending duel message")
			websocket.Message.Send(ws, core.MessageDuel)
		}

		if len(nearestUsersLocations) == 0 {
			log.Printf("Error no near users")
			continue
		}

		for _, nearestUserLocation := range nearestUsersLocations {
			log.Printf("Nearest user is %d for user %d", nearestUserLocation.UserID, userID)
		}
	}
}
