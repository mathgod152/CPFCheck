package handlers 

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mathgod152/CFPcheck/internal/dto"
)

func (_r *Router) ValidateCnpjHandler(c *fiber.Ctx) error {
	var cnpj dto.CnpjValidatorDTO
	if err := c.BodyParser(&cnpj); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload" + fmt.Sprint(err),
		})
	}
	validateFormat, err := _r.CnpjValidator.ConvertToValidateFormat(cnpj.Cnpj)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error 500",
		})
	}
	result := _r.CnpjValidator.CnpjValidatorEntity.Verify(validateFormat)
	if !result {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"response": result,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": result,
	})
}