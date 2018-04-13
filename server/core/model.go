package core

import "time"

// Model is the generic type for all models
type Model struct {
	ID        uint       `json:"id" form:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt" form:"createdAt" gorm:"not null;type:timestamp;default:now()"`
	UpdatedAt time.Time  `json:"updatedAt" form:"updatedAt" gorm:"not null;type:timestamp;default:now()"`
	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt" gorm:"type:date"`
}
