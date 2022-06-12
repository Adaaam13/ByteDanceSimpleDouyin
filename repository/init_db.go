package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open("mysql", "root:Arsenal!13cql@(127.0.0.1:3306)/bddb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Video{}, &User{}, &Comment{}, &Favorite{}, &Follow{})
}