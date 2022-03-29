package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kahono922/chitchat/database"
)

func addStudent(c *fiber.Ctx) error {
	type Student struct {
		Name     string
		Reg      string
		Password string
	}
	student := new(Student)

	if err := c.BodyParser(student); err != nil {
		return c.JSON(fiber.Map{
			"err":    err,
			"status": "FAIL",
		})
	}
	db := database.DB
	err := db.Create(&student).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "ERROR",
			"err":    err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   "Items Added Successfully",
	})
}
func getStudent(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	type Student struct {
		Name     string
		Reg      string
		Password string
	}
	var student Student
	if err := db.Where("id = ?", id).First(&student).Error; err != nil {
		return c.JSON(fiber.Map{
			"status": "ERROR",
			"err":    err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   student,
	})
}
func updateStudent(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	val := c.Params("type")
	value := c.Params("val")

	if val != "name" && val != "reg" && val != "password" {
		return c.JSON(fiber.Map{
			"status": "ERROR",
			"err":    "Invalid Type",
		})
	}
	if err:=db.Exec("UPDATE students SET "+val+" = ? WHERE id = ?", value, id).Error; err != nil {
		return c.JSON(fiber.Map{
			"status": "ERROR",
			"err":    err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   "Items Updated Successfully",
	})
	
}
