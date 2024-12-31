package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mathgod152/CFPcheck/internal/dto"
)

func (_r *Router) ValidateCpfHandler(c *fiber.Ctx) error {
	var cpf dto.CpfValidatorDTO
	if err := c.BodyParser(&cpf); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload" + fmt.Sprint(err),
		})
	}
	validateFormat, err := _r.CpfValidator.ConvertToValidateFormat(cpf.CPF)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error 500",
		})
	}
	result := _r.CpfValidator.CpfValidatorEntity.Verify(validateFormat)
	if !result {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"response": result,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": result,
	})
}