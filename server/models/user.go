package models

import "github.com/makeorbreak-io/shooting-stars/server/core"

// User is the model for users
type User struct {
	core.Model

	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	PasswordHash string `json:"password" form:"password"`
	Gender       string `json:"gender" form:"gender"`
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

// TableName returns the table name for the model
func (User) TableName() string {
	return "mob_users"
}
