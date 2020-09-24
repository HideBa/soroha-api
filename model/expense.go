package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Expense struct {
	gorm.Model
	Price    int       `gorm:"not null"`
	UsedDate time.Time `gorm:"not null"`
	Comment  string
	User     User
	UserID   uint
}
