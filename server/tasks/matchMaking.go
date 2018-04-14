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
	activeUsers, err := task.LocationService.GetActiveUsers()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	for _, user := range activeUsers {
		log.Printf("Active user: %v", user)
	}
}
