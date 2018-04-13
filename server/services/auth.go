package services

import (
	"crypto/rand"
	"encoding/base64"
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

// GenerateAuthToken generates an authentication token for a given user ID
func (service *AuthService) GenerateAuthToken(userID uint) (*models.AuthToken, error) {
	token, err := service.generateRandomString(32)
	if err != nil {
		return nil, err
	}

	authToken := models.AuthToken{
		Token:  token,
		UserID: userID,
	}

	db := service.Database.Create(&authToken)
	if db.Error != nil {
		return nil, db.Error
	}

	return &authToken, nil
}

// GenerateRandomBytes returns securely generated random bytes.
func (service *AuthService) generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded securely generated random string.
func (service *AuthService) generateRandomString(s int) (string, error) {
	b, err := service.generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

// ValidateLogin checks if the password matches the user hashed password
func (service *AuthService) ValidateLogin(password string, user *models.User) error {
	return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
}
