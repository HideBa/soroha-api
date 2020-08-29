package request

import (
	"github.com/HideBa/soroha-api/model"
	"github.com/labstack/echo"
)

type UserRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"passwordd" validate:"required"`
	} `json:"user"`
}

func (userRequest *UserRegisterRequest) Bind(c echo.Context, user *model.User) error {
	if err := c.Bind(userRequest); err != nil {
		return err
	}

	if err := c.Validate(userRequest); err != nil {
		return err
	}
	user.Username = userRequest.User.Username
	h, err := user.HashPassword(userRequest.User.Password)
	if err != nil {
		return err
	}
	user.Password = h
	return nil
}
