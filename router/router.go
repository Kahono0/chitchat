package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpRoutes(app *fiber.App) {
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	api := app.Group("/api")
	food:=api.Group("/food")
	student:=api.Group("/student")

	food.Get("/:id",getFood)

	food.Post("/",addFood)
	
	food.Put("/:id/:type/:val",updateFood)
	student.Post("/",addStudent)
	student.Get("/:id",getStudent)
	student.Put("/:id/:type/:val",updateStudent)
	api.Post("/order",completeOrder)
	api.Post("/wallet/:id",addWallet)

}
