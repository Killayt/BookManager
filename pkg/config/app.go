package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	d, err := gorm.Open("mysql", "killlayt:js#oa23X)x!@12@/simplerest?charset=utf8&parseTime=True&loc=Local") // Paste your db info here
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
