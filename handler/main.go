package handler

import (
	"net/http"

	"github.com/HideBa/soroha-api/user"
	"github.com/labstack/echo"
)

type Handler struct {
	userStore user.Store
}

func (h *Handler) MainPage(c echo.Context) error {
	// db.Init()
	// return func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello")
	// }
	return c.String(http.StatusOK, "Hello")
}

func NewHandler(us user.Store) *Handler {
	return &Handler{
		userStore: us,
	}
}
