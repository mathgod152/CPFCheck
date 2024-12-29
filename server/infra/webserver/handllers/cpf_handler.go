package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mathgod152/CFPcheck/internal/dto"
)

func (_r *Router) CreateCpfHandler(c *fiber.Ctx) error {
	var cpf dto.CpfDTO
	if err := c.BodyParser(&cpf); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}
	validateFormat, err := _r.Cpf.NewCpf(cpf)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error 500" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": validateFormat,
	})
}