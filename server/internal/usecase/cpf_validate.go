package usecase

import (
	"fmt"

	"github.com/mathgod152/CFPcheck/internal/entity"
)

type CpfValidatorUseCase struct {
	CpfValidatorEntity entity.CpfValidatorInterface
}

func (c *CpfValidatorUseCase) ConvertToValidateFormat(cpfNumber string)([]int, error)  {
	verify, err := c.CpfValidatorEntity.ConverteToIntArray(cpfNumber)
	if err != nil{
		return nil, fmt.Errorf("Erro To Convert CPF Number: %v", err)
	}
	return verify, nil
}

func (c *CpfValidatorUseCase) CpfValidate(cpfNumber []int) bool {
	verify := c.CpfValidatorEntity.Verify(cpfNumber)

	return verify
}
