package models

import (
	"github.com/aodin/date"
	"github.com/makeorbreak-io/shooting-stars/server/core"
)

// LoginRequest holds the login request information
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest holds the registration request information
type RegisterRequest struct {
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirmPassword"`
	BirthDate       date.Date `json:"birthDate"`
	Gender          string    `json:"gender"`
}

// AuthToken is the model for an authentication token
type AuthToken struct {
	core.Model

	Token  string `json:"token"`
	UserID uint   `json:"userID"`
}

// IAuthService is the service for authentication
type IAuthService interface {
	CreateTable() error
	GenerateAuthToken(userID uint) (*AuthToken, error)
	ValidateLogin(password string, user *User) error
}
