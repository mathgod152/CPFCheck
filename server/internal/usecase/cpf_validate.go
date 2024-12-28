package usecase

import (

	"github.com/mathgod152/CFPcheck/internal/entity"
)

type CpfValidatorUseCase struct {
	CpfValidatorEntity entity.CpfValidatorInterface
}

func (c *CpfValidatorUseCase) ConvertToValidateFormat(cpfNumber []int) (bool) {
	verify := c.CpfValidatorEntity.Verify(cpfNumber)
 
	 return verify;
 }

func (c *CpfValidatorUseCase) CpfValidate(cpfNumber []int) (bool) {
   verify := c.CpfValidatorEntity.Verify(cpfNumber)

	return verify;
}


