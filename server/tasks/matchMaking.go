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
var connections = make(map[uint]*websocket.Conn, 0)

// AddConnection adds a web socket connection associated to a given user
func (task *MatchMakingTask) AddConnection(userID uint, ws *websocket.Conn) {
	connections[userID] = ws
}

// RemoveConnection removes a web socket connection associated to a given user
func (task *MatchMakingTask) RemoveConnection(userID uint) {
	_, ok := connections[userID]
	if ok {
		delete(connections, userID)
	}
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
		return
	}

	for _, userID := range activeUsersIDs {
		userLocation, err := task.LocationService.Get(userID)
		if err != nil {
			log.Printf("Error when getting current user locations: %v", err)
			continue
		}

		// Check if in match
		if match, _ := task.MatchService.GetActiveMatchByUserID(userID, core.GetConfiguration().MatchTimeout); match != nil {
			continue
		}

		// Nearest users
		nearestUsersLocations, err := task.LocationService.GetNearestActiveUserLocation(userID, core.GetConfiguration().MaxLocationLastUpdate)
		if err != nil {
			log.Printf("Error when getting nearest active users locations: %v", err)
			continue
		}
		if len(nearestUsersLocations) == 0 {
			continue
		}

		// Match with someone
		for _, nearestUserLocation := range nearestUsersLocations {
			log.Printf("We have near users...")
			if match, _ := task.MatchService.GetActiveMatchByUserID(nearestUserLocation.ID, core.GetConfiguration().MatchTimeout); match != nil {
				continue
			}

			log.Printf("That are not in match")

			// Check distance
			if distance := nearestUserLocation.Location.GeoDistanceFrom(&userLocation.Location, false); distance > core.GetConfiguration().MaxDistance {
				continue
			}

			log.Printf("Near than %v", core.GetConfiguration().MaxDistance)

			// Create match
			match := &models.Match{
				StartTime: time.Now(),
				UserOneID: userID,
				UserTwoID: nearestUserLocation.UserID,
			}
			_, err = task.MatchService.Create(match)
			if err != nil {
				log.Printf("Error while creating the match: %v", err)
				continue
			}

			log.Printf("Match created")

			// Sending the duel
			if ws, exists := connections[userID]; exists {
				websocket.Message.Send(ws, core.MessageDuel)
			}

			log.Printf("Message sent")

			break
		}
	}
}
