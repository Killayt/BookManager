package config

import (
	"os"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	d, err := gorm.Open("mysql", "killayt:"+os.Getenv("BOOK_MANAGER_PASSWORD_DB"+"?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
