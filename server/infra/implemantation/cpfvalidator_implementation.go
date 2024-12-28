package implemantation

import (
	"fmt"

	"github.com/mathgod152/CFPcheck/internal/entity"
)

var _ entity.CpfValidatorInterface = (*CpfValidatorImplementation)(nil)

type CpfValidatorImplementation struct{}

func (c *CpfValidatorImplementation) ConverteToIntArray(cpf string) ([]int, error) {
	return nil, nil
}


func (c *CpfValidatorImplementation) Verify(cpfNumber []int) (bool) {
	if len(cpfNumber) != 11 {
        return false
    }
	if allElementsEqual(cpfNumber){
		return false
	}
	if !verifyFirstDigit(cpfNumber){
		return false
	}
	if !verifySecondDigit(cpfNumber){
		return false
	}
    fmt.Println("CPF válido!")
    return true
}

func verifyFirstDigit(cpfNumber []int)(bool){
	firstNine := cpfNumber[:9]
	firstVerifierDigit := cpfNumber[9]
	var sumFirstVerifierDigit int
    for i, num := range firstNine {
        sumFirstVerifierDigit += num * (10 - i)
    }
    sumFirstVerifierDigit = (sumFirstVerifierDigit * 10) % 11
    if sumFirstVerifierDigit >= 10 {
        sumFirstVerifierDigit = 0
    }

    if firstVerifierDigit != sumFirstVerifierDigit {
        fmt.Printf("Erro no primeiro dígito: esperado %v, calculado %v\n", firstVerifierDigit, sumFirstVerifierDigit)
        return false
    }
	return true
}

func verifySecondDigit(cpfNumber []int)(bool){
	secondVerifierDigit := cpfNumber[10]
    firstTen := cpfNumber[:10]
    var sumSecondVerifierDigit int
    for i, num := range firstTen {
        sumSecondVerifierDigit += num * (11 - i)
    }
    sumSecondVerifierDigit = (sumSecondVerifierDigit * 10) % 11
    if sumSecondVerifierDigit >= 10 {
        sumSecondVerifierDigit = 0
    }

    if secondVerifierDigit != sumSecondVerifierDigit {
        fmt.Printf("Erro no segundo dígito: esperado %v, calculado %v\n", secondVerifierDigit, sumSecondVerifierDigit)
        return false
    }
	return true
}

func allElementsEqual[T comparable](arr []T) bool {
	if len(arr) == 0 {
		return true // Considera vazio como "todos iguais"
	}

	first := arr[0]
	for _, v := range arr {
		if v != first {
			return false
		}
	}
	return true
}