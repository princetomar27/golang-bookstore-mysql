package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// The purpose of this app.go is to return variable DB, to interact with the mysql database

var (
	db * gorm.DB
)

func ConnectToDatabase(){

	// open connection with the db, using gorm 
	database, err := gorm.Open("mysql", "root:prince123@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")


	if err != nil{
		panic(err)
	}

	db = database
} 

func GetDB() *gorm.DB{
	return db
}