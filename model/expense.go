package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Expense struct {
	gorm.Model
	Slug  uuid.UUID `gorm:"unique_index;not null"`
	Price int       `gorm:"not null"`
	// UsedDate time.Time `gorm:"not null"`
	Comment string
	User    User
	UserID  uint `gorm:"not null"`
}