package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Calculation struct {
	gorm.Model
	CalculatedAt time.Time `gorm:"not null" json:"calculatedAt"`
	Price        int       `gorm:"not null" json:"price"`
	IsPaid       bool      `gorm:"not null" json:"isPaid"`
	Users        []User    `gorm:"many2many:user_calculations"`
	Team         Team      `gorm:"not null"`
}
