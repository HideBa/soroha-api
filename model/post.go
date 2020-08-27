package model

import "github.com/jinzhu/gorm"

type (
	Post struct {
		gorm.Model
		Text string `gorm:"not null" json:"text"`
		User User   `gorm:"foreignkey:user_id" json:"user"`
		UID  uint   `gorm:"column:user_id"`
	}
)
