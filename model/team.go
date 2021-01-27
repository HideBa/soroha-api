package model

import "github.com/jinzhu/gorm"

type (
	Team struct {
		gorm.Model
		TeamName     string `gorm:"unique_index;not null" json:"teamName"`
		IsPersonal   bool   `json:"isPersonal"`
		Users        []User `gorm:"many2many:user_teams" json:"name"`
		Calculations []Calculation
	}
)
