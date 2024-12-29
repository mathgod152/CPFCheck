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
	CpfValidateImplementation *implemantation.CpfValidatorImplementation
)

func TestCpfValidate(t *testing.T) {
	f := fuzz.New()
	cpfValidateUseCase := &usecase.CpfValidatorUseCase{
		CpfValidatorEntity: CpfValidateImplementation, // Injetando a implementação 
	}
	for i := 0; i < 1000; i++ {
		newTrueCpf := generateTrueCpf(f)
		cpfCheck := cpfValidateUseCase.CpfValidate(newTrueCpf)

		assert.Equal(t, cpfCheck, true)
	}
	newFakeCpf := generateFakeCpf(f)

	for i := 0; i < 1000; i++ {
		cpfCheck := cpfValidateUseCase.CpfValidate(newFakeCpf)
		assert.NotEqual(t, cpfCheck, true)
	}
}

func TestConvertToValidateFormatWithFuzzer(t *testing.T) {
	cpfValidateUseCase := &usecase.CpfValidatorUseCase{
		CpfValidatorEntity: CpfValidateImplementation, // Injetando a implementação 
	}

	f := fuzz.New()

	// Testar CPFs válidos
	for i := 0; i < 1000; i++ {
		trueCpf := generateTrueCpf(f)
		cpfStr := convertIntArrayToString(trueCpf)

		t.Run(fmt.Sprintf("Valid CPF #%d", i), func(t *testing.T) {
			result, err := cpfValidateUseCase.ConvertToValidateFormat(cpfStr)
			assert.NoError(t, err)
			assert.Equal(t, trueCpf, result)
		})
	}

	// Testar CPFs inválidos
	for i := 0; i < 1000; i++ {
		fakeCpf := generateFakeCpf(f)
		cpfStr := convertIntArrayToString(fakeCpf)

		t.Run(fmt.Sprintf("Invalid CPF #%d", i), func(t *testing.T) {
			result, err := cpfValidateUseCase.ConvertToValidateFormat(cpfStr)
			assert.NoError(t, err)
			assert.Equal(t, fakeCpf, result)
		})
	}
}

func generateTrueCpf(f *fuzz.Fuzzer) []int {
	firstNine := generateFirstNineNumbers(f)
	cpf := append(firstNine, generateCpfFirstVerifierDigit(firstNine))
	cpf = append(cpf, generateCpfSecondVerifierDigit(cpf))
	return cpf
}
func generateFakeCpf(f *fuzz.Fuzzer) []int {
	firstNine := generateFirstNineNumbers(f)
	cpf := append(firstNine, generateCpfFakeFirstVerifierDigit(firstNine))
	cpf = append(cpf, generateCpfFakeSecondVerifierDigit(cpf))
	return cpf
}

func generateFirstNineNumbers(f *fuzz.Fuzzer) []int {
	var cpfFirstNine []int
	for len(cpfFirstNine) < 9 {
		var newNumber int
		f.Fuzz(&newNumber)
		newNumber = int(math.Abs(float64(newNumber))) % 10
		cpfFirstNine = append(cpfFirstNine, newNumber)
	}
	if allElementsEqual(cpfFirstNine) {
		cpfFirstNine = generateFirstNineNumbers(f)
	}
	return cpfFirstNine
}

func generateCpfFirstVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit += num * (10 - i)
	}
	verifyDigit = (verifyDigit * 10) % 11
	if verifyDigit == 10 {
		verifyDigit = 0
	}
	return verifyDigit
}

func generateCpfSecondVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit += num * (11 - i)
	}
	verifyDigit = (verifyDigit * 10) % 11
	if verifyDigit == 10 {
		verifyDigit = 0
	}
	return verifyDigit
}
func generateCpfFakeFirstVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (10 - i))
	}
	verifyDigit = (verifyDigit * 10) % 11
	if verifyDigit >= 9 {
		verifyDigit = 1
	}
	return verifyDigit + 1
}

func generateCpfFakeSecondVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (11 - i))
	}
	verifyDigit = (verifyDigit * 10) % 11
	if verifyDigit >= 9 {
		verifyDigit = 1
	}
	return verifyDigit + 1
}

func convertIntArrayToString(cpf []int) string {
	cpfStr := ""
	for i, num := range cpf {
		cpfStr += fmt.Sprintf("%d", num)
		if i == 2 || i == 5 {
			cpfStr += "."
		} else if i == 8 {
			cpfStr += "-"
		}
	}
	return cpfStr
}

func allElementsEqual[T comparable](arr []T) bool {
	if len(arr) == 0 {
		return true 
	}

	first := arr[0]
	for _, v := range arr {
		if v != first {
			return false
		}
	}
	return true
}
