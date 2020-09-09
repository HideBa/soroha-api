package handler

import (
	"github.com/labstack/echo"
)

func (h *Handler) Register(v1 *echo.Group) {
	// jwtMiddleware := router.JWT(config.GetConfig())
	v1.GET("", h.MainPage)
	guestUsers := v1.Group("/users")
	guestUsers.POST("/signup", h.SignUp)
	guestUsers.POST("/signin", h.Login)

	// user := v1.Group("user", jwtMiddleware)
	// fmt.Println(user)
	// user.GET("", h.CurrentUser)
	// user.PATCH("", h.UpdateUser)
	user := v1.Group("/user")
	user.POST("/expenses", h.CreateExpense)
}
