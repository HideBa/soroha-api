package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not nul" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Token    string `json:"token"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Token    string `json:"token,omitempty"`
	CreateAt string `json:"created_at"`
}

type UserTweetsLikesResponse struct {
	User string `json:"user"`
	// Tweets []TweetResponse `json:"tweets"`
}

func (user *User) UserTransformer() UserResponse {
	return UserResponse{
		ID:       user.ID,
		UserName: user.Username,
		CreateAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
