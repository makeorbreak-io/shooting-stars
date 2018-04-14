package models

import (
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/paulmach/go.geo"
)

// UpdateLocationRequest is the model for a location update request
type UpdateLocationRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Location is the model for locations
type Location struct {
	core.Model

	UserID   uint      `json:"userID" gorm:"unique"`
	Location geo.Point `json:"location" gorm:"type:geography"`
}

// ILocationService is the service for locations
type ILocationService interface {
	CreateTable() error
	Get(userID uint) (*Location, error)
	Create(userID uint, latitude, longitude float64) (uint, error)
	Update(userID uint, latitude, longitude float64) error
}