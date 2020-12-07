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
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB.DbUser, cfg.DB.DbPass, cfg.DB.DbHost, cfg.DB.DbPort, cfg.DB.DbName)
	// url := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB.DbUser, cfg.DB.DbPass, cfg.DB.DbHost, cfg.DB.DbName)
	fmt.Println("url", url)
	// db, err = gorm.Open("mysql", url)
	db, err = gorm.Open("mysql", config.GetConfig().DB.DbURL)
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
		&model.Expense{},
		&model.Team{},
		&model.Calculation{},
	)
	db.Model(&model.Expense{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
