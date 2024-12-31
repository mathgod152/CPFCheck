package midware

import (
	"github.com/mathgod152/CFPcheck/infra/implemantation"
)

func ValidateCpf(cpf string) (bool){
	imp := implemantation.CpfValidatorImplementation{}
	if len(cpf) != 11  &&  len(cpf) != 14 {
		return false
	}
	cpfToValidate, err := imp.ConverteToIntArray(cpf)
	if err != nil {
		return false
	}
	validate := imp.Verify(cpfToValidate)
	if !validate{
		return false
	}
	return true
}