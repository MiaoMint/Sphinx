package database

import (
	"log"

	"github.com/MiaoMint/Sphinx/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"))
	if err != nil {
		log.Fatalln(err)
	}

	err = DB.AutoMigrate(&model.Domain{}, &model.API{}, &model.Log{})
	if err != nil {
		log.Fatalln(err)
	}
}
