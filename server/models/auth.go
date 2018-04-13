package models

import "github.com/makeorbreak-io/shooting-stars/server/core"

// LoginRequest holds the login request information
type LoginRequest struct {
	Email    string
	Password string
}

// RegisterRequest holds the registration request information
type RegisterRequest struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword"`
	Gender          string `json:"gender" form:"gender"`
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
	GenerateAuthToken(userID uint) *AuthToken
	ValidateLogin(password string, user *User) error
}
