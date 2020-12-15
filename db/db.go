package db

import (
	"log"

	"github.com/HideBa/soroha-api/config"
	"github.com/HideBa/soroha-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
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
