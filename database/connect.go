package database

import (
	"fmt"

	"github.com/kahono922/chitchat/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "moringa"
	port     = "5432"
	password = "Access"
)

var DB *gorm.DB
var err error

func Conn(){

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, "chitchat")
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s", host, port, user, password)
		DB, err = gorm.Open(postgres.Open(dsn))

		if err != nil {
			panic(err)
		}
		db := DB.Exec("CREATE DATABASE chitchat")
		if db.Error != nil {
			panic(db.Error)
		}
	}
	DB.AutoMigrate(models.Food{})
	DB.AutoMigrate(models.Student{})
}
