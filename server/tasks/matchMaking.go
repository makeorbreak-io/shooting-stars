package tasks

import (
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"log"
)

// Ensure MatchMakingTask implements ITask.
var _ ITask = &MatchMakingTask{}

// MatchMakingTask is the task related to the match making
type MatchMakingTask struct {
	LocationService models.ILocationService
	MatchService    models.IMatchService
}

// Run executes the match making task
func (task *MatchMakingTask) Run() {
	activeUsersIDs, err := task.LocationService.GetActiveUsers(99999)
	if err != nil {
		log.Printf("Error when getting active users: %v", err)
		return
	}

	for _, userID := range activeUsersIDs {
		nearestUserLocation, err := task.LocationService.GetNearestUserLocation(userID)
		if err != nil {
			log.Printf("Error when getting nearest user location: %v", err)
			continue
		}

		log.Printf("Nearest user is %d for user %d", nearestUserLocation.UserID, userID)
	}
}
