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

// IUserService is the service for users
type IUserService interface {
	CreateTable() error
	Get(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) (uint, error)
	Update(user *User) error
	Delete(id uint) error
}
