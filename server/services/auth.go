package services

import (
	"github.com/jinzhu/gorm"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"golang.org/x/crypto/bcrypt"
)

// Ensure AuthService implements IAuthService.
var _ models.IAuthService = &AuthService{}

// AuthService is a struct that holds methods to handle authentication
type AuthService struct {
	Database *gorm.DB
}

// CreateTable creates the users table
func (service *AuthService) CreateTable() error {
	err := service.Database.DropTableIfExists(&models.AuthToken{}).Error
	if err != nil {
		return err
	}

	return service.Database.CreateTable(&models.AuthToken{}).Error
}

// Get returns a user by its ID
func (service *AuthService) GenerateAuthToken(userID uint) *models.AuthToken {
	authToken := models.AuthToken{
		Token:  "",
		UserID: userID,
	}

	return &authToken
}

// ValidateLogin checks if the password matches the user hashed password
func (service *AuthService) ValidateLogin(password string, user *models.User) error {
	return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
}
