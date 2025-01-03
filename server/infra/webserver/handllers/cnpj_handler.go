package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mathgod152/CFPcheck/infra/webserver/midware"
	"github.com/mathgod152/CFPcheck/internal/dto"
)

func (_r *Router) CreateCnpjHandler(c *fiber.Ctx) error {
	var cnpj dto.CnpjDTO
	if err := c.BodyParser(&cnpj); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}
	if  !midware.ValidateCnpj(cnpj.CnpjNumber){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cnpj INVALIDO",
		})
	}
	validateFormat, err := _r.Cnpj.NewCnpj(cnpj)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error 500" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": validateFormat,
	})
}

func (_r *Router) GetAllCnpjsHandler(c *fiber.Ctx) error {
	cnpjs, err := _r.Cnpj.SelectCnpjs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving Cnpjs" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": cnpjs,
	})
}

func (_r *Router) GetCnpjHandler(c *fiber.Ctx) error {
	cnpjNumber := c.Params("cnpj")
	if strings.Contains(cnpjNumber, "%2F"){
		cnpjNumber = strings.ReplaceAll(cnpjNumber, "%2F", "/")
	}
	cnpj, err := _r.Cnpj.SelectById(cnpjNumber)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cnpj not found" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": cnpj,
	})
}

func (_r *Router) UpdateCnpjHandler(c *fiber.Ctx) error {
	cnpjNumber := c.Params("cnpj")
	if strings.Contains(cnpjNumber, "%2F"){
		cnpjNumber = strings.ReplaceAll(cnpjNumber, "%2F", "/")
	}
	fmt.Println("Cnpj: ", cnpjNumber)
	var cnpj dto.CnpjDTO
	if err := c.BodyParser(&cnpj); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}
	_, err := _r.Cnpj.UpdateCnpj(cnpj, cnpjNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating Cnpj" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": "Atualizado com suceesso",
	})
}

func (_r *Router) DeleteCnpjHandler(c *fiber.Ctx) error {
	cnpjNumber := c.Params("cnpj")
	if strings.Contains(cnpjNumber, "%2F"){
		cnpjNumber = strings.ReplaceAll(cnpjNumber, "%2F", "/")
	}
	success, err := _r.Cnpj.DeleteCnpj(cnpjNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting Cnpj" + fmt.Sprint(err),
		})
	}
	if !success {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cnpj not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cnpj deleted successfully",
	})
}

func (_r *Router) AddToBlockListCnpjHandler(c *fiber.Ctx) error {
	cnpjNumber := c.Params("cnpj")
	if strings.Contains(cnpjNumber, "%2F"){
		cnpjNumber = strings.ReplaceAll(cnpjNumber, "%2F", "/")
	}
	success, err := _r.Cnpj.AddCnpjToBlockList(cnpjNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting Cnpj" + fmt.Sprint(err),
		})
	}
	if !success {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cnpj not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cnpj Was added Black List",
	})
}

func (_r *Router) GetBlocklistCnpjsHandler(c *fiber.Ctx) error {
	cnpjs, err := _r.Cnpj.SelectBlockListCnpjs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving Cnpjs" + fmt.Sprint(err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": cnpjs,
	})
}

func (_r *Router) RemoveToBlockListCnpjHandler(c *fiber.Ctx) error {
	cnpjNumber := c.Params("cnpj")
	if strings.Contains(cnpjNumber, "%2F"){
		cnpjNumber = strings.ReplaceAll(cnpjNumber, "%2F", "/")
	}
	success, err := _r.Cnpj.RemoveCnpjFromBlockList(cnpjNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting Cnpj" + fmt.Sprint(err),
		})
	}
	if !success {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cnpj not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cnpj Was Reemove Black List",
	})
}


