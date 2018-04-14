package tasks

import (
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
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

// Start is a function to start the task with a given interval between runs
func (task *MatchMakingTask) Start(interval uint) {
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				task.Run()
			case <-quit:
				ticker.Stop()
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

	for _, userID := range activeUsersIDs {
		nearestUsersLocations, err := task.LocationService.GetNearestUserLocation(userID)
		if err != nil {
			log.Printf("Error when getting nearest user location: %v", err)
			continue
		}

		if len(nearestUsersLocations) == 0 {
			log.Printf("No users near %d", userID)
			continue
		}

		for _, nearestUserLocation := range nearestUsersLocations {
			log.Printf("Nearest user is %d for user %d", nearestUserLocation.UserID, userID)
		}
	}
}
