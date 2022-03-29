package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kahono922/chitchat/database"
)

type Food struct {
	Name      string
	Price     float32
	Available bool
}

func addFood(c *fiber.Ctx) error {
	food := new(Food)

	if err := c.BodyParser(food); err != nil {
		return c.JSON(fiber.Map{
			"err":    err,
			"status": "FAIL",
		})
	}
	db := database.DB
	err := db.Create(&food).Error
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

func getFood(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	type Food struct {
		Name      string
		Price     float32
		Available bool
	}
	var food Food
	if err := db.Where("id = ?", id).First(&food).Error; err != nil {
		return c.JSON(fiber.Map{
			"status": "ERROR",
			"err":    err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   food,
	})
}

func updateFood(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	val := c.Params("type")
	value := c.Params("val")
	if val != "price" && val != "available" {
		return c.JSON(fiber.Map{
			"status": "ERROR",
			"err":    "Invalid type",
		})
	}

	if err:=db.Exec("UPDATE foods SET "+val+" = ? WHERE id = ?", value, id).Error; err != nil {
		return c.JSON(fiber.Map{
			"status": "ERROR",
			"err":    err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   "Item Updated Successfully",
	})
}

