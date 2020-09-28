package model

import "github.com/jinzhu/gorm"

type (
	Team struct {
		gorm.Model
		TeamName string `gorm:"not null" json:"teamname"`
		Users    []User `gorm:"many2many:user_teams" json:"name"`
	}
)
