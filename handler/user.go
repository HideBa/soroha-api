package handler

import (
	"net/http"

	"github.com/HideBa/soroha-api/model"
	"github.com/HideBa/soroha-api/request"
	"github.com/HideBa/soroha-api/response"
	util "github.com/HideBa/soroha-api/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUp(c echo.Context) error {
	var user model.User
	req := &request.UserRegisterRequest{}
	if err := req.Bind(c, &user); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error}
	}
	if err := h.userStore.Create(&user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	return c.JSON(http.StatusCreated, response.NewUserResponse(&user))
}

func (h *Handler) Login(c echo.Context) error {
	req := &request.UserLoginRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	u, err := h.userStore.GetByUsername(req.User.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, util.AccessForbidden())
	}
	if !u.CheckPassword(req.User.Password) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "failure to authenticate"}
	}

	return c.JSON(http.StatusOK, response.NewUserResponse(u))
}

func (h *Handler) CurrentUser(c echo.Context) error {
	u, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, util.NotFound())
	}

	return c.JSON(http.StatusOK, response.NewUserResponse(u))
}

func userIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}

func (h *Handler) CreateTeam(c echo.Context) error {
	var teamModel model.Team
	req := &request.TeamCreateRequest{}
	if err := req.Bind(c, &teamModel); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	if err := h.userStore.CreateTeam(&teamModel, userIDFromToken(c)); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	return c.JSON(http.StatusOK, response.NewTeamResponse(&teamModel))
}

func (h *Handler) TeamsList(c echo.Context) error {
	var (
		teams []model.Team
		err   error
	)

	userID := userIDFromToken(c)
	teams, err = h.userStore.TeamsList(userID, 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, response.TeamsListResponse(h.userStore, userID, teams))
}

func (h *Handler) TeamUsersList(c echo.Context) error {
	var (
		team  model.Team
		users []model.User
		err   error
	)

	teamName := c.Param("name")

	team, users, err = h.userStore.TeamUsersList(teamName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, response.TeamUsersListResponse(&team, users))
}

func (h *Handler) AddUserOnTeam(c echo.Context) error {
	var usersNames []string
	teamName := c.Param("name")

	team, err := h.userStore.TeamByName(teamName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.NewError(err))
	}
	if team == nil {
		return c.JSON(http.StatusNotFound, util.NotFound())
	}
	req := &request.AddRemoveUsersRequest{}
	usersNames, err = req.Bind(c, usersNames)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	userModels, err := h.userStore.AddUserOnTeam(team, usersNames)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}

	return c.JSON(http.StatusOK, response.TeamUsersListResponse(team, userModels))
}
