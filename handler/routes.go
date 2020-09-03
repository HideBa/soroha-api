package handler

import "github.com/labstack/echo"

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := midleware
	v1.GET("", h.MainPage)

	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)
}
