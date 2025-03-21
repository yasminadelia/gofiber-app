package main

import (
	database "gofiber-app/config"
	"gofiber-app/migrations"
	"gofiber-app/models"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
		_ = godotenv.Load()

		port := os.Getenv("PORT")
		if port == "" {
			port = "3000" // Default port
		}
		
    app := fiber.New()
		
		// Connect to DB
		database.ConnectDB()

		 // Run Migrations
		if err := migrations.Migrate(database.DB); err != nil {
			panic("Migration failed: " + err.Error())
		}

		// Auto-migrate models
		// database.DB.AutoMigrate(&models.User{})

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Fiber! 🚀")
    })

		app.Post("/users", func(c *fiber.Ctx) error {
				type CreateUserInput struct {
						Name  string `json:"name"`
						Email string `json:"email"`
				}
		
				var input CreateUserInput
		
				if err := c.BodyParser(&input); err != nil {
						return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
				}
		
				user := models.User{
						Name:  input.Name,
						Email: input.Email,
				}
		
				database.DB.Create(&user)
		
				return c.JSON(user)
		})
	

		app.Post("/user", func(c *fiber.Ctx) error {
			type User struct {
				Name string `json:"name"`
				Age int `json:"age"`
			}

			var user User

			if err :=c.BodyParser(&user); err != nil {
				return c.Status(400).SendString(("Invalid JSON"))
			}

			return c.JSON(fiber.Map{
				"message": "User created!",
				"user": user,
			})
		})

		log.Fatal(app.Listen(":" + port))
		log.Println("Server running on port", port)

	
}

