package response

import (
	"github.com/HideBa/soroha-api/model"
	util "github.com/HideBa/soroha-api/utils"
)

type UserResponse struct {
	User struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	} `json:"user"`
}

func NewUserResponse(userModel *model.User) *UserResponse {
	res := &UserResponse{}
	res.User.Username = userModel.Username
	res.User.Token = util.GenerateJWT(userModel.ID)
	return res
}
