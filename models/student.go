package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Id       int `gorm:"primary_key"`
	Name     string
	Reg      string `gorm:"unique"`
	Password string
}

type Wallet struct {
	gorm.Model
	Id        int `gorm:"primary_key"`
	Balance   float32
	StudentId string
}
