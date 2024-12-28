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

func TestCheck(t *testing.T) {
	f := fuzz.New()
	cpfUseCase := &usecase.CpfValidatorUseCase{
		CpfValidatorEntity: CpfValidateImplementation, // Injetando a implementação 
	}
	for i := 0; i < 1000; i++ {
		newTrueCpf := generateTrueCpf(f)
		cpfCheck := cpfUseCase.CpfValidate(newTrueCpf)

		assert.Equal(t, cpfCheck, true)
	}
	newFakeCpf := generateFakeCpf(f)

	for i := 0; i < 1000; i++ {
		cpfCheck := cpfUseCase.CpfValidate(newFakeCpf)

		assert.NotEqual(t, cpfCheck, true)
	}
}

func generateTrueCpf(f *fuzz.Fuzzer) []int {
	firstNine := generateFirstNineNumbers(f)
	cpf := append(firstNine, generateCpfFirstVerifierDigit(firstNine))
	cpf = append(cpf, generateCpfSecondVerifierDigit(cpf))
	fmt.Println("Cpf True: ", cpf)
	return cpf
}
func generateFakeCpf(f *fuzz.Fuzzer) []int {
	firstNine := generateFirstNineNumbers(f)
	cpf := append(firstNine, generateCpfFakeFirstVerifierDigit(firstNine))
	cpf = append(cpf, generateCpfFakeSecondVerifierDigit(cpf))
	fmt.Println("Cpf Fake: ", cpf)
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
		fmt.Println("I: ", i, "Num: ", num)
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
		fmt.Println("I: ", i, "Num: ", num)
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
		fmt.Println("I: ", i, "Num: ", num)
		verifyDigit = verifyDigit + (num * (10 - i))
	}
	verifyDigit = (verifyDigit * 10) % 11
	if verifyDigit > 9 {
		return verifyDigit - 1
	}
	return verifyDigit + 1
}

func generateCpfFakeSecondVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		fmt.Println("I: ", i, "Num: ", num)
		verifyDigit = verifyDigit + (num * (11 - i))
	}
	verifyDigit = (verifyDigit * 10) % 11
	if verifyDigit > 9 {
		return verifyDigit - 1
	}
	return verifyDigit + 1
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
