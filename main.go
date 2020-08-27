package main

import (
	"github.com/HideBa/soroha-api/config"
	"github.com/HideBa/soroha-api/db"
	"github.com/HideBa/soroha-api/handler"

	// "soroha-api/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	// userGroup := e.Group("/user")

	config := config.GetConfig()
	db.Init()
	dbm := db.GetDB()
	log.Print("-----------will try to migrate")
	h := &handler.Handler{DB: dbm}
	db.AutoMigrate(dbm)
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.Server.KEY),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/login" || c.Path() == "/signup" || c.Path() == "/" {
				return true
			}
			return false
		},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", h.MainPage)
	e.POST("/signup", h.SignUp)

	// e.Logger.Fatal(e.Start(":" + config.Server.PORT))
	e.Logger.Fatal(e.Start(":3000"))
}
