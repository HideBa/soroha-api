package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	Handler struct {
		DB *gorm.DB
	}
)

func (h *Handler) MainPage(c echo.Context) error {
	// db.Init()
	// return func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello")
	// }
	return c.String(http.StatusOK, "Hello")
}

// func MainPage() echo.HandlerFunc {
// 	db.Init()
// 	return func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello")
// 	}
// }

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		DB: db,
	}
}
