package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Id        int `gorm:"primary_key"`
	Name      string
	Price     float32
	Available bool
}
