package models

import (
	"github.com/makeorbreak-io/shooting-stars/server/core"
)

// Location is the model for locations
type Location struct {
	core.Model

	UserID    uint    `json:"userID" gorm:"unique"`
	Latitude  float32 `json:"latitude" gorm:"type:"`
	Longitude float32 `json:"longitude"`
}

// ILocationService is the service for locations
type ILocationService interface {
	CreateTable() error
	Get(userID uint) (*Location, error)
	Create(userID uint, latitude, longitude float32) (uint, error)
	Update(userID uint, latitude, longitude float32) error
}
