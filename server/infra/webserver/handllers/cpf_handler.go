package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mathgod152/CFPcheck/infra/webserver/midware"
	"github.com/mathgod152/CFPcheck/internal/dto"
)

func (_r *Router) CreateCpfHandler(c *fiber.Ctx) error {
	var cpf dto.CpfDTO
	if err := c.BodyParser(&cpf); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}
	if  !midware.ValidateCpf(cpf.CpfNumber){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "CPF INVALIDO",
		})
	}
	_, err := _r.Cpf.NewCpf(cpf)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error 500" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": "Atualizado com suceesso",
	})
}

func (_r *Router) GetAllCpfsHandler(c *fiber.Ctx) error {
	cpfs, err := _r.Cpf.SelectCpfs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving CPFs" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": cpfs,
	})
}

func (_r *Router) GetCpfHandler(c *fiber.Ctx) error {
	cpfNumber := c.Params("cpf")
	cpf, err := _r.Cpf.SelectById(cpfNumber)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "CPF not found" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": cpf,
	})
}

func (_r *Router) UpdateCpfHandler(c *fiber.Ctx) error {
	cpfNumber := c.Params("cpf")
	fmt.Println("CPF: ", cpfNumber)
	var cpf dto.CpfDTO
	if err := c.BodyParser(&cpf); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}
	updatedCpf, err := _r.Cpf.UpdateCpf(cpf, cpfNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating CPF" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": updatedCpf,
	})
}

func (_r *Router) DeleteCpfHandler(c *fiber.Ctx) error {
	cpfNumber := c.Params("cpf")
	success, err := _r.Cpf.DeleteCpf(cpfNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting CPF" + fmt.Sprint(err),
		})
	}
	if !success {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "CPF not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CPF deleted successfully",
	})
}


func (_r *Router) AddToBlockListCpfHandler(c *fiber.Ctx) error {
	cpfNumber := c.Params("cpf")
	success, err := _r.Cpf.AddCpfToBlockList(cpfNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting CPF" + fmt.Sprint(err),
		})
	}
	if !success {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "CPF not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CPF Was added Black List",
	})
}

func (_r *Router) GetBlocklistCpfsHandler(c *fiber.Ctx) error {
	cpfs, err := _r.Cpf.SelectBlockListCpfs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving CPFs" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": cpfs,
	})
}

func (_r *Router) RemoveToBlockListCpfHandler(c *fiber.Ctx) error {
	cpfNumber := c.Params("cpf")
	success, err := _r.Cpf.RemoveCpfFromBlockList(cpfNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting CPF" + fmt.Sprint(err),
		})
	}
	if !success {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "CPF not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "CPF Was Reemove Black List",
	})
}