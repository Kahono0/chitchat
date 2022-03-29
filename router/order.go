package router

import(
	"github.com/kahono922/chitchat/database"
	"github.com/kahono922/chitchat/models"
	"github.com/gofiber/fiber/v2"

)

func completeOrder(c *fiber.Ctx) error{
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.JSON(fiber.Map{
			"err":    err,
			"status": "FAIL",
		})
	}
	//confirm items
	db := database.DB
	var total float32
	for _, v := range  order.Food{
		//check if item is available
		food := new(models.Food)
		if err:=db.Where("id = ?", v.Id).First(food).Error; err != nil {
			return c.JSON(fiber.Map{
				"status": "ERROR",
				"err":    err,
			})
		}
		if !food.Available{
			return c.JSON(fiber.Map{
				"status": "ERROR",
				"err":    "Item not available",
			})
		}
		total+=food.Price
		
	}
	//check if total is enough
	if total > getBalance(order.StudentId){
		return c.JSON(fiber.Map{
			"status": "ERROR",
			"err":    "Not enough balance",
		})
	}
	//update balance
	if _,err:=updateBalance(order.StudentId, total); err != nil {
		return c.JSON(fiber.Map{
			"status": "ERROR",
			"err":    err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   "Order Completed Successfully",
	})

}