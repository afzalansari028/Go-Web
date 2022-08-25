package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:king@/library?charset=utf8&parseTime=True&loc=Local")
	// d, err := gorm.Open("mysql", "root:king@tcp(127.0.0.1:3306)/library")
	if err != nil {
		panic(err)
	}
	db = d
	log.Print(d)
}

func GetDB() *gorm.DB {
	return db
}
