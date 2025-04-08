package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

func FormatValidationError(err error) fiber.Map {
	errors := make(map[string]string)

	// Handle validator errors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			fieldName := fieldErr.Field()

			var msg string
			switch fieldErr.Tag() {
			case "required":
				msg = fieldName + " is required"
			case "email":
				msg = "Invalid email format"
			case "min":
				msg = fieldName + " must be at least " + fieldErr.Param() + " characters"
			case "max":
				msg = fieldName + " must be at most " + fieldErr.Param() + " characters"
			default:
				msg = fieldName + " is invalid"
			}

			errors[fieldErr.Field()] = msg
		}

		return fiber.Map{"errors": errors}
	}

	// Handle DB errors (e.g., duplicate key)
	if strings.Contains(err.Error(), "duplicate key") {
		return fiber.Map{
			"error": "A record with that value already exists.",
		}
	}

	// Generic fallback
	return fiber.Map{
		"error": err.Error(),
	}

}