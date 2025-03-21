package routes

import (
	"gofiber-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber! 🚀")
	})

	userGroup := app.Group("/users")

	userGroup.Post("/", controllers.CreateUser)
	userGroup.Get("/", controllers.GetUsers)
	userGroup.Get("/:id", controllers.GetUser)
	userGroup.Put("/:id", controllers.UpdateUser)
	userGroup.Delete("/:id", controllers.DeleteUser)
}
