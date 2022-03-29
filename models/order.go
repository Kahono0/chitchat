package models

type OrderedFood struct {
	Id string
	Name string
	Quantity int
}

type Order struct {
	Id        int `gorm:"primary_key"`
	StudentId string
	StudentPass string
	Food []OrderedFood
	
}

