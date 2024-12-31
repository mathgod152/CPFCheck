package midware

import (
	"github.com/mathgod152/CFPcheck/infra/implemantation"
)

func ValidateCnpj(cnpj string) (bool){
	imp := implemantation.CnpjValidatorImplementation{}
	if len(cnpj) != 14  &&  len(cnpj) != 18 {
		return false
	}
	cnpjToValidate, err := imp.ConverteToIntArray(cnpj)
	if err != nil {
		return false
	}
	validate := imp.Verify(cnpjToValidate)
	if !validate{
		return false
	}
	return true
}