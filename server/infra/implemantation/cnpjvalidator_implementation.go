package implemantation

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/mathgod152/CFPcheck/internal/entity"
)

var _ entity.CnpjValidatorInterface = (*CnpjValidatorImplementation)(nil)

type CnpjValidatorImplementation struct{}

func (c *CnpjValidatorImplementation) ConverteToIntArray(cnpj string) ([]int, error) {
	if len(cnpj) != 14  &&  len(cnpj) != 18 {
		return nil, errors.New("Cnpj Invalido")
	}
	// Remove caracteres não numéricos
	cnpjDigits := regexp.MustCompile(`\D`).ReplaceAllString(cnpj, "")


	intArray := make([]int, len(cnpjDigits))
	for i, char := range cnpjDigits {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, errors.New("erro ao converter CNPJ para inteiros")
		}
		intArray[i] = num
	}

	return intArray, nil
}

func (c *CnpjValidatorImplementation) Verify(cnpjNumber []int) bool {
	if len(cnpjNumber) != 14 {
		return false
	}
	if allElementsEqual(cnpjNumber) {
		return false
	}
	if !verifyFirstCnpjDigit(cnpjNumber) {
		return false
	}
	if !verifySecondCnpjDigit(cnpjNumber) {
		return false
	}
	return true
}

func verifyFirstCnpjDigit(cnpjNumber []int) bool {
	firstTwelve := cnpjNumber[:12]
	firstVerifierDigit := cnpjNumber[12]

	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var sum int
	for i, num := range firstTwelve {
		sum += num * weights[i]
	}

	sum = 11 - (sum % 11)
	if sum >= 10 {
		sum = 0
	}

	return firstVerifierDigit == sum
}

func verifySecondCnpjDigit(cnpjNumber []int) bool {
	firstThirteen := cnpjNumber[:13]
	secondVerifierDigit := cnpjNumber[13]

	weights := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var sum int
	for i, num := range firstThirteen {
		sum += num * weights[i]
	}

	sum = 11 - (sum % 11)
	if sum >= 10 {
		sum = 0
	}

	return secondVerifierDigit == sum
}

