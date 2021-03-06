package handler

import (
	"github.com/HideBa/soroha-api/config"
	"github.com/HideBa/soroha-api/router/middleware"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT([]byte(config.GetConfig().Server.KEY))
	v1.GET("", h.MainPage)
	guestUsers := v1.Group("/users")
	guestUsers.POST("/signup", h.SignUp)
	guestUsers.POST("/signin", h.Login)
	guestUsers.GET("", h.List)

	user := v1.Group("/user", jwtMiddleware)
	user.GET("", h.CurrentUser)
	team := v1.Group("/teams", jwtMiddleware)
	team.POST("", h.CreateTeam)
	team.GET("", h.TeamsList)
	team.GET("/:name/users", h.TeamUsersList)
	team.POST("/:teamname/users", h.AddUserOnTeam)
	team.GET("/:teamname/user/expenses", h.UserExpenses)
	team.GET("/:teamname/expenses", h.TeamExpenses)
	team.POST("/:teamname/calculations", h.CalculateExpenses)
	team.PATCH("/:teamname/calculations/:slug", h.UpdateCalculation)
	team.GET("/:name/calculations", h.Calculations)
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
	expenses.PATCH("/:slug", h.UpdateExpense)
	expenses.DELETE("/:slug", h.DeleteExpense)

	// calculations := v1.Group("/calculations", jwtMiddleware)

	images := v1.Group("/file", jwtMiddleware)
	images.POST("", h.SignHandler)
}
