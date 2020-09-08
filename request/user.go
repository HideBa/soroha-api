package request

import (
	"github.com/HideBa/soroha-api/model"
	"github.com/labstack/echo"
)

type UserRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (req *UserRegisterRequest) Bind(c echo.Context, user *model.User) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	// TODO: must validate later
	// if err := c.Validate(req); err != nil {
	// 	return err
	// }
	user.Username = req.User.Username
	h, err := user.HashPassword(req.User.Password)
	if err != nil {
		return err
	}
	user.Password = h
	return nil
}

type UserLoginRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (req *UserLoginRequest) Bind(c echo.Context) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	// if err := c.Validate(req); err != nil {
	// 	return err
	// }
	return nil
}
