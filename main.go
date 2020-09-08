package main

import (
	"github.com/HideBa/soroha-api/config"
	"github.com/HideBa/soroha-api/db"
	"github.com/HideBa/soroha-api/handler"
	"github.com/HideBa/soroha-api/store"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	apiV1 := e.Group("/api/v1")

	config := config.GetConfig()
	db.Init()
	dbm := db.GetDB()
	db.AutoMigrate(dbm)

	userStore := store.NewUserStore(dbm)
	h := handler.NewHandler(userStore)
	e.Logger.SetLevel(log.ERROR)
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(config.Server.KEY),
	// 	Skipper: func(c echo.Context) bool {
	// 		if c.Path() == "api/v1/users/login" || c.Path() == "api/v1/users" || c.Path() == "api/v1" {
	// 			return true
	// 		}
	// 		return false
	// 	},
	// }))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	h.Register(apiV1)
	e.Logger.Fatal(e.Start(":" + config.Server.PORT))
}
