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
	apiV1 := e.Group("/api/v1")

	config := config.GetConfig()
	fmt.Println("---", config)
	db.Init()
	dbm := db.GetDB()
	log.Print("-----------will try to migrate")
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

	h.Register(apiV1)
	// e.POST("/signup", h.SignUp)
	// e.Logger.Fatal(e.Start(":" + config.Server.PORT))
	e.Logger.Fatal(e.Start(":3000"))
}
