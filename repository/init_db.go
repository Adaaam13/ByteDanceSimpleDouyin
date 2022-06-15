package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	dsn := "root:Arsenal!13cql@(127.0.0.1:3306)/bddb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &Video{}, &Comment{}, &Favorite{}, &Follow{})
}
