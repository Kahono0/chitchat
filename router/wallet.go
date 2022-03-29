package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kahono922/chitchat/database"
	"github.com/kahono922/chitchat/models"

)

func getBalance(id string) float32{
	db := database.DB
	var balance float32
	if err:=db.Where("id = ?", id).First(&balance).Error; err != nil {
		return -1
	}
	return balance
}

func updateBalance(id string, amount float32) (float32, error) {
	db := database.DB
	var wallet models.Wallet
	var balance float32

	//begin transaction
	tx := db.Begin()
	if err := tx.Where("id = ?", id).First(&wallet).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	balance = wallet.Balance - amount
	if err := tx.Model(&wallet).Update("balance", balance).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return balance, nil
}

func addWallet(c *fiber.Ctx) error{
	db := database.DB
	wallet := new(models.Wallet)
	if err := c.BodyParser(wallet); err != nil {
		return c.JSON(fiber.Map{
			"err":    err,
			"status": "FAIL",
		})
	}
	err := db.Create(&wallet).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "ERROR",
			"err":    err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   "Wallet Added Successfully",
	})

}