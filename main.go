package main

import (
	database "gofiber-app/config"
	"gofiber-app/migrations"
	"gofiber-app/routes"
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

		routes.UserRoutes(app)


		log.Fatal(app.Listen(":" + port))
		log.Println("Server running on port", port)

	
}

