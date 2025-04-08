package controllers

import (
	database "gofiber-app/config"
	"gofiber-app/models"
	"gofiber-app/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	type CreateUserInput struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email" gorm:"unique"`
		Age int `json:"age"`
		Address string `json:"address"`
	}

	var input CreateUserInput

	if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := utils.Validate.Struct(input); err != nil {
		return c.Status(400).JSON(utils.FormatValidationError(err))
	}

	user := models.User{
			Name:  input.Name,
			Email: input.Email,
			Age: input.Age,
			Address: input.Address,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(utils.FormatValidationError(err))
	}

	return c.JSON(user)
	
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)

	return c.JSON(users)
}


func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	type UpdateUserInput struct {
		Name  string `json:"name" `
    Email string `json:"email" gorm:"unique"`
		Age int `json:"age"`
		Address string `json:"address"`
	}

	var input UpdateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := utils.Validate.Struct(input); err != nil {
		return c.Status(400).JSON(utils.FormatValidationError(err))
	}

	if err := database.DB.Model(&user).Updates(input).Error; err != nil {
		return c.Status(400).JSON(utils.FormatValidationError(err))
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	database.DB.Delete(&user)

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}