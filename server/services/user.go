package services

import (
	"github.com/jinzhu/gorm"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
)

// Ensure UserService implements IUserService.
var _ models.IUserService = &UserService{}

// UserService is a struct that holds methods to handle users
type UserService struct {
	Database *gorm.DB
}

// CreateTable creates the users table
func (service *UserService) CreateTable() error {
	err := service.Database.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}

	return service.Database.CreateTable(&models.User{}).Error
}

// Get returns a user by its ID
func (service *UserService) Get(id uint) (*models.User, error) {
	var result models.User
	db := service.Database.First(&result, id)

	if db.Error != nil {
		return nil, db.Error
	}

	return &result, nil
}

// GetByEmail returns a user by its email
func (service *UserService) GetByEmail(email string) (*models.User, error) {
	var result models.User
	db := service.Database.First(&result, "email = ?", email)

	if db.Error != nil {
		return nil, db.Error
	}

	return &result, nil
}

// GetAll returns all users
func (service *UserService) GetAll() ([]*models.User, error) {
	var result []*models.User
	db := service.Database.Find(&result)

	if db.Error != nil {
		return nil, db.Error
	}

	return result, nil
}

// GetUsersMostWins returns the users with most wins
func (service *UserService) GetUsersMostWins(limit uint) ([]*models.User, error) {
	var result []*models.User
	db := service.Database.Order("wins DESC").Limit(limit).Find(&result)

	if db.Error != nil {
		return nil, db.Error
	}

	return result, nil
}

// AddWin adds a win to a user
func (service *UserService) AddWin(userID uint) error {
	user, err := service.Get(userID)
	if err != nil {
		return err
	}

	user.Wins += 1

	return service.Update(user)
}

// Create saves a new user in the database
func (service *UserService) Create(user *models.User) (uint, error) {
	if !service.Database.NewRecord(user) {
		return 0, core.ErrorDuplicated
	}

	db := service.Database.Create(user)
	if db.Error != nil {
		return 0, db.Error
	}

	return user.ID, nil
}

// Update updates the information about a user in the database
func (service *UserService) Update(user *models.User) error {
	return service.Database.Save(user).Error
}

// Delete soft deletes an user from the database
func (service *UserService) Delete(id uint) error {
	user, err := service.Get(id)
	if err != nil {
		return err
	}

	return service.Database.Delete(user).Error
}
