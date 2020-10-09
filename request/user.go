package request

import (
	"github.com/HideBa/soroha-api/model"
	"github.com/labstack/echo/v4"
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

type TeamCreateRequest struct {
	Team struct {
		TeamName string `json:"teamName" validate:"required"`
	} `json:"team"`
}

func (req *TeamCreateRequest) Bind(c echo.Context, team *model.Team) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	team.TeamName = req.Team.TeamName
	return nil
}

type AddRemoveUsersRequest struct {
	Users []struct {
		UserName string `json:"userName"`
	} `json:"users"`
}

func (req *AddRemoveUsersRequest) Bind(c echo.Context, users []string) ([]string, error) {
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	for _, user := range req.Users {
		users = append(users, user.UserName)
	}
	return users, nil
}
