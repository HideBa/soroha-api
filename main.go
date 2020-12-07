package main

import (
	"github.com/HideBa/soroha-api/config"
	"github.com/HideBa/soroha-api/db"
	"github.com/HideBa/soroha-api/handler"
	"github.com/HideBa/soroha-api/router"
	"github.com/HideBa/soroha-api/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	expenseStore := store.NewExpenseStore(dbm)
	h := handler.NewHandler(userStore, expenseStore)
	e.Logger.SetLevel(log.ERROR)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Validator = router.NewValidator()

	h.Register(apiV1)
	e.Logger.Fatal(e.Start(":" + config.Server.PORT))
}
