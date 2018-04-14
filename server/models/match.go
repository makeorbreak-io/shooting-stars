package models

import (
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"time"
)

// Match is the model for matches
type Match struct {
	core.Model

	StartTime        time.Time  `json:"startTime"`
	UserOneID        uint       `json:"userOneID"`
	UserTwoID        uint       `json:"userTwoID"`
	UserOneShootTime *time.Time `json:"userOneShootTime" gorm:"null"`
	UserTwoShootTime *time.Time `json:"userTwoShootTime" gorm:"null"`
	WinnerID         *uint       `json:"winnerID"`
}

// IMatchService is the service for matches
type IMatchService interface {
	CreateTable() error
	Get(id uint) (*Match, error)
	GetActiveMatchByUserID(userID, timeout uint) (*Match, error)
	GetAll() ([]*Match, error)
	Create(match *Match) (uint, error)
	Update(match *Match) error
	Delete(id uint) error
}
