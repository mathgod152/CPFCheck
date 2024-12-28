package usecase

import (

	"github.com/mathgod152/CFPcheck/internal/entity"
)

type CpfValidatorUseCase struct {
	CpfValidatorEntity entity.CpfValidatorInterface
}

func (c *CpfValidatorUseCase) CpfValidate(cpfNumber []int) (bool) {
   varify := c.CpfValidatorEntity.Verify(cpfNumber)
}


