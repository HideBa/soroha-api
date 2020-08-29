package response

import "github.com/HideBa/soroha-api/model"

type UserResponse struct {
	User struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	} `json:"user"`
}

func NewUserResponse(userModel *model.User) *UserResponse {
	r := &UserResponse{}
	r.User.Username = userModel.Username
	r.User.Token = userModel.Token
}