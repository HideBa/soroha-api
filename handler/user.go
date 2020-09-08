package handler

import (
	"net/http"

	"github.com/HideBa/soroha-api/model"
	"github.com/HideBa/soroha-api/request"
	"github.com/HideBa/soroha-api/response"
	util "github.com/HideBa/soroha-api/utils"
	"github.com/labstack/echo"
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
		return c.JSON(http.StatusForbidden, util.NewError(err))
	}
	if !u.CheckPassword(req.User.Password) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "failure to authenticate"}
	}

	return c.JSON(http.StatusOK, response.NewUserResponse(u))
}
