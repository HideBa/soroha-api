package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Calculation struct {
	gorm.Model
	Slug         uuid.UUID `gorm:"unique_index;not null"`
	CalculatedAt time.Time `gorm:"not null" json:"calculatedAt"`
	Price        int       `gorm:"not null" json:"price"`
	IsPaid       bool      `gorm:"not null" json:"isPaid"`
	UserID       uint      `gorm:"not null"`
	User         User      `gorm:"not null`
	TeamID       uint      `gorm:"not null"`
	Team         Team      `gorm:"not null"`
}
