package services

import (
	"github.com/jinzhu/gorm"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"github.com/paulmach/go.geo"
	"log"
)

// Ensure LocationService implements ILocationService.
var _ models.ILocationService = &LocationService{}

// LocationService is a struct that holds methods to handle locations
type LocationService struct {
	Database *gorm.DB
}

// CreateTable creates the locations table
func (service *LocationService) CreateTable() error {
	err := service.Database.DropTableIfExists(&models.Location{}).Error
	if err != nil {
		return err
	}

	return service.Database.CreateTable(&models.Location{}).Error
}

// Get returns a location by its user ID
func (service *LocationService) Get(userID uint) (*models.Location, error) {
	var result models.Location
	db := service.Database.First(&result, "user_id = ?", userID)

	if db.Error != nil {
		return nil, db.Error
	}

	log.Printf("Location: %v", result)

	return &result, nil
}

// Create saves a new match in the database
func (service *LocationService) Create(userID uint, latitude, longitude float64) (uint, error) {
	location := models.Location{
		UserID:     userID,
		LocationDB: geo.NewPointFromLatLng(latitude, longitude).ToWKT(),
	}

	db := service.Database.Create(&location)
	if db.Error != nil {
		return 0, db.Error
	}

	return location.ID, nil
}

// Update updates the information about a match in the database
func (service *LocationService) Update(userID uint, latitude, longitude float64) error {
	return service.Database.Model(&models.Location{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{"location": geo.NewPointFromLatLng(latitude, longitude).ToWKT()}).Error
}
