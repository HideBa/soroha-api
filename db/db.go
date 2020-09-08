package db

import (
	"fmt"
	"log"

	"github.com/HideBa/soroha-api/config"
	"github.com/HideBa/soroha-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	cfg := config.GetConfig()
	// TODO:enable to use container DB
	url := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB.DbUser, cfg.DB.DbPass, cfg.DB.DbHost, cfg.DB.DbName)
	fmt.Print("--", url)
	db, err = gorm.Open("mysql", url)
	if err != nil {
		log.Fatal("failured to connect with db")
	}
}

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
	// db.Model(&model.Post{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
