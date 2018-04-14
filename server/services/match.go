package services

import (
	"github.com/jinzhu/gorm"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
)

// Ensure MatchService implements IMatchService.
var _ models.IMatchService = &MatchService{}

// MatchService is a struct that holds methods to handle matches
type MatchService struct {
	Database *gorm.DB
}

// CreateTable creates the matches table
func (service *MatchService) CreateTable() error {
	err := service.Database.DropTableIfExists(&models.Match{}).Error
	if err != nil {
		return err
	}

	return service.Database.CreateTable(&models.Match{}).Error
}

// Get returns a match by its ID
func (service *MatchService) Get(id uint) (*models.Match, error) {
	var result models.Match
	db := service.Database.First(&result, id)

	if db.Error != nil {
		return nil, db.Error
	}

	return &result, nil
}

// GetAll returns all matches
func (service *MatchService) GetAll() ([]*models.Match, error) {
	var result []*models.Match
	db := service.Database.Find(&result)

	if db.Error != nil {
		return nil, db.Error
	}

	return result, nil
}

// Create saves a new match in the database
func (service *MatchService) Create(match *models.Match) (uint, error) {
	if !service.Database.NewRecord(match) {
		return 0, core.ErrorDuplicated
	}

	db := service.Database.Create(match)
	if db.Error != nil {
		return 0, db.Error
	}

	return match.ID, nil
}

// Update updates the information about a match in the database
func (service *MatchService) Update(match *models.Match) error {
	return service.Database.Save(match).Error
}

// Delete soft deletes an match from the database
func (service *MatchService) Delete(id uint) error {
	match, err := service.Get(id)
	if err != nil {
		return err
	}

	return service.Database.Delete(match).Error
}
