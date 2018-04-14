package services

import (
	"github.com/jinzhu/gorm"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"github.com/paulmach/go.geo"
	"time"
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

	// Get location as a geo.Point
	type Location struct {
		Location geo.Point `gorm:"location"`
	}
	location := Location{}
	db = service.Database.Table("mob_locations").
		Select("ST_AsBinary(location) as location").
		Where("user_id = ?", userID).
		Scan(&location)
	if db.Error != nil {
		return nil, db.Error
	}

	result.Location = location.Location

	return &result, nil
}

// GetActiveUsers returns the list of active users
func (service *LocationService) GetActiveUsers(maxLastUpdate uint) ([]uint, error) {
	var result []uint
	db := service.Database.Model(&models.Location{}).
		Where("updated_at >= ?", time.Now().Add(-1*time.Second*time.Duration(maxLastUpdate))).Pluck("user_id", &result)

	if db.Error != nil {
		return nil, db.Error
	}

	return result, nil
}

// GetNearestUserLocation returns the nearest user location to a given user
func (service *LocationService) GetNearestUserLocation(userID uint) (*models.Location, error) {
	userLocation, err := service.Get(userID)
	if err != nil {
		return nil, err
	}

	var result models.Location
	db := service.Database.
		Order("ST_Distance(location, '" + userLocation.Location.ToWKT() + "')").
		Where("user_id != ?", userID).
		First(&result)

	if db.Error != nil {
		return nil, db.Error
	}

	return &result, nil
}

// Create saves a new location in the database
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

// Update updates the information about a location in the database
func (service *LocationService) Update(userID uint, latitude, longitude float64) error {
	return service.Database.Model(&models.Location{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{"location": geo.NewPointFromLatLng(latitude, longitude).ToWKT()}).Error
}
