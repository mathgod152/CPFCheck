package usecase

import (
	"fmt"

	"github.com/mathgod152/CFPcheck/internal/entity"
)

type CpfUseCase struct {
	CpfUseCase entity.CpfInterface
}

func (c *CpfUseCase) CpfValidate(cpfNumber []int) (bool, error) {
    if len(cpfNumber) != 11 {
        return false, fmt.Errorf("CPF must contain exactly 11 digits")
    }
	if allElementsEqual(cpfNumber){
		return false, fmt.Errorf("CPF cannot have all numbers the same.")
	}

    firstVerifierDigit := cpfNumber[9]
    secondVerifierDigit := cpfNumber[10]
    firstNine := cpfNumber[:9]
	fmt.Printf("CPF recebido: %v\n", cpfNumber)
    fmt.Printf("Primeiros 9 dígitos: %v\n", firstNine)
    // Cálculo do primeiro dígito verificador
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
        return false, nil
    }
	
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
        return false, nil
    }

    fmt.Println("CPF válido!")
    return true, nil
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
