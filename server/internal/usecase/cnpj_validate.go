package usecase

import (
	"fmt"

	"github.com/mathgod152/CFPcheck/internal/entity"
)

type CnpjValidatorUseCase struct {
	CnpjValidatorEntity entity.CnpjValidatorInterface
}

func (c *CnpjValidatorUseCase) ConvertToValidateFormat(cnpjNumber string)([]int, error)  {
	verify, err := c.CnpjValidatorEntity.ConverteToIntArray(cnpjNumber)
	if err != nil{
		return nil, fmt.Errorf("Erro To Convert Cnpj Number: %v", err)
	}
	return verify, nil
}

func (c *CnpjValidatorUseCase) CnpjValidate(cnpjNumber []int) bool {
	verify := c.CnpjValidatorEntity.Verify(cnpjNumber)

	return verify
}
