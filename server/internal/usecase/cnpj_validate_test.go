package usecase_test

import (
	"fmt"
	"math"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/mathgod152/CFPcheck/infra/implemantation"
	"github.com/mathgod152/CFPcheck/internal/usecase"
	"github.com/stretchr/testify/assert"
)

var (
	CnpjValidateImplementation *implemantation.CnpjValidatorImplementation
)

func TestCnpjValidate(t *testing.T) {
	f := fuzz.New()
	cnpjValidateUseCase := &usecase.CnpjValidatorUseCase{
		CnpjValidatorEntity: CnpjValidateImplementation, // Injetando a implementação 
	}

	// Testar CNPJs válidos
	for i := 0; i < 1000; i++ {
		newTrueCnpj := generateTrueCnpj(f)
		cnpjCheck := cnpjValidateUseCase.CnpjValidate(newTrueCnpj)

		assert.Equal(t, cnpjCheck, true)
	}
}

func TestConvertCnpjToValidateFormatWithFuzzer(t *testing.T) {
	cnpjValidateUseCase := &usecase.CnpjValidatorUseCase{
		CnpjValidatorEntity: CnpjValidateImplementation, // Injetando a implementação 
	}

	f := fuzz.New()

	// Testar CNPJs válidos
	for i := 0; i < 1000; i++ {
		trueCnpj := generateTrueCnpj(f)
		cnpjStr := convertIntArrayToString(trueCnpj)

		t.Run(fmt.Sprintf("Valid CNPJ #%d", i), func(t *testing.T) {
			result, err := cnpjValidateUseCase.ConvertToValidateFormat(cnpjStr)
			assert.NoError(t, err)
			assert.Equal(t, trueCnpj, result)
		})
	}

	// Testar CNPJs inválidos
	for i := 0; i < 1000; i++ {
		fakeCnpj := generateFakeCnpj(f)
		cnpjStr := convertIntArrayToString(fakeCnpj)

		t.Run(fmt.Sprintf("Invalid CNPJ #%d", i), func(t *testing.T) {
			result, err := cnpjValidateUseCase.ConvertToValidateFormat(cnpjStr)
			assert.NoError(t, err)
			assert.Equal(t, fakeCnpj, result)
		})
	}
}

func generateTrueCnpj(f *fuzz.Fuzzer) []int {
	firstTwelve := generateFirstTwelveNumbers(f)
	cnpj := append(firstTwelve, generateCnpjFirstVerifierDigit(firstTwelve))
	cnpj = append(cnpj, generateCnpjSecondVerifierDigit(cnpj))
	fmt.Println("CNPJ",cnpj)
	return cnpj
}

func generateFakeCnpj(f *fuzz.Fuzzer) []int {
	firstTwelve := generateFirstTwelveNumbers(f)
	cnpj := append(firstTwelve, generateCnpjFakeFirstVerifierDigit(firstTwelve))
	cnpj = append(cnpj, generateCnpjFakeSecondVerifierDigit(cnpj))
	return cnpj
}

func generateFirstTwelveNumbers(f *fuzz.Fuzzer) []int {
	var cnpjFirstTwelve []int
	for len(cnpjFirstTwelve) < 12 {
		var newNumber int
		f.Fuzz(&newNumber)
		newNumber = int(math.Abs(float64(newNumber))) % 10
		cnpjFirstTwelve = append(cnpjFirstTwelve, newNumber)
	}
	if allElementsEqual(cnpjFirstTwelve) {
		cnpjFirstTwelve = generateFirstTwelveNumbers(f)
	}
	return cnpjFirstTwelve
}

func generateCnpjFirstVerifierDigit(fd []int) int {
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var sum int
	for i, num := range fd {
		sum += num * weights[i]
	}
	verifyDigit := 11 - (sum % 11)
	if verifyDigit >= 10 {
		verifyDigit = 0
	}
	return verifyDigit
}

func generateCnpjSecondVerifierDigit(fd []int) int {
	weights := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var sum int
	for i, num := range fd {
		sum += num * weights[i]
	}
	verifyDigit := 11 - (sum % 11)
	if verifyDigit >= 10 {
		verifyDigit = 0
	}
	return verifyDigit
}

func generateCnpjFakeFirstVerifierDigit(fd []int) int {
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var sum int
	for i, num := range fd {
		sum += num * weights[i]
	}
	verifyDigit := (11 - (sum % 11)) + 1
	if verifyDigit >= 10 {
		verifyDigit = 1
	}
	return verifyDigit
}

func generateCnpjFakeSecondVerifierDigit(fd []int) int {
	weights := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var sum int
	for i, num := range fd {
		sum += num * weights[i]
	}
	verifyDigit := (11 - (sum % 11)) + 1
	if verifyDigit >= 10 {
		verifyDigit = 1
	}
	return verifyDigit
}


