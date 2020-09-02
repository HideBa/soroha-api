package main

import (
	"fmt"

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
	// userGroup := e.Group("/user")

	config := config.GetConfig()
	fmt.Println("---", config)
	db.Init()
	dbm := db.GetDB()
	log.Print("-----------will try to migrate")
	db.AutoMigrate(dbm)

	userStore := store.NewUserStore(dbm)
	h := handler.NewHandler(userStore)
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
