package handler

import (
	"fmt"

	"github.com/HideBa/soroha-api/config"
	"github.com/HideBa/soroha-api/router/middleware"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(config.GetConfig())
	fmt.Println(jwtMiddleware)
	v1.GET("", h.MainPage)
	guestUsers := v1.Group("/users")
	guestUsers.POST("/signup", h.SignUp)
	guestUsers.POST("/signin", h.Login)

	// user := v1.Group("user", jwtMiddleware)
	// fmt.Println(user)
	// user.GET("", h.CurrentUser)
	// user.PATCH("", h.UpdateUser)
	// user := v1.Group("/user")

	// expenses := v1.Group("/expenses", middleware.JWTConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(config.GetConfig().Server.KEY),
	// 	Skipper: func(c echo.Context) bool {
	// 		if c.Request().Method == "GET" {
	// 			return true
	// 		}
	// 		return false
	// 	},
	// }))
	expenses := v1.Group("/expenses", middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			if c.Request().Method == "GET" {
				return true
			}
			return false
		}, SigningKey: []byte(config.GetConfig().Server.KEY),
	}))
	expenses.POST("", h.CreateExpense)
	expenses.GET("", h.Expenses)
	expenses.PUT("/:slug", h.UpdateExpense)
	expenses.DELETE("/:slug", h.DeleteExpense)
}
