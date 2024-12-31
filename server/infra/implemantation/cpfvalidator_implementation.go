package implemantation

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/mathgod152/CFPcheck/internal/entity"
)


var _ entity.CpfValidatorInterface = (*CpfValidatorImplementation)(nil)
 
type CpfValidatorImplementation struct{}

func (c *CpfValidatorImplementation) ConverteToIntArray(cpf string) ([]int, error) {
	if len(cpf) != 11  &&  len(cpf) != 14 {
		return nil, errors.New("CPF Invalido")
	}
	cpfDigits := regexp.MustCompile(`\D`).ReplaceAllString(cpf, "")

	if len(cpfDigits) != 11 {
		return nil, errors.New("o CPF deve conter exatamente 11 dígitos")
	}

	intArray := make([]int, len(cpfDigits))
	for i, char := range cpfDigits {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, errors.New("erro ao converter CPF para inteiros")
		}
		intArray[i] = num
	}

	return intArray, nil
}


func (c *CpfValidatorImplementation) Verify(cpfNumber []int) (bool) {
	if len(cpfNumber) != 11 {
        return false
    }
	if allElementsEqual(cpfNumber){
		return false
	}
	if !verifyCPFFirstDigit(cpfNumber){
		return false
	}
	if !verifyCPFSecondDigit(cpfNumber){
		return false
	}
    return true
}

func verifyCPFFirstDigit(cpfNumber []int)(bool){
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
        return false
    }
	return true
}

func verifyCPFSecondDigit(cpfNumber []int)(bool){
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