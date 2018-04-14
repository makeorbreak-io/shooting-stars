package models

import (
	"github.com/aodin/date"
	"github.com/makeorbreak-io/shooting-stars/server/core"
)

// User is the model for users
type User struct {
	core.Model

	Name         string    `json:"name"`
	Email        string    `json:"email" gorm:"unique"`
	PasswordHash string    `json:"password"`
	BirthDate    date.Date `json:"birthDate"`
	Gender       string    `json:"gender"`
	Wins         uint      `json:"wins" gorm:"default:0"`
}

// UserRank is the model that holds a rank of a user
type UserRank struct {
	Name   string `json:"name"`
	Wins   uint   `json:"wins"`
	UserID uint   `json:"userID"`
	Rank   uint   `json:"rank"`
}

// IUserService is the service for users
type IUserService interface {
	CreateTable() error
	Get(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]*User, error)
	GetUsersMostWins(limit uint) ([]*User, error)
	Create(user *User) (uint, error)
	Update(user *User) error
	Delete(id uint) error
}
