package usecase_test

import (
	"fmt"
	"math/rand"
	"testing"

	//fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	//f := fuzz.New()
	newTrueCpf := generateTrueCpf()

	cpfCheck, err := CpfCheckUseCaseMock.UsecaseCpf.Check(*&newTrueCpf)

	assert.Nil(t, err)
	assert.Equal(t, cpfCheck, true)

	newFakeCpf := generateFakeCpf()

	cpfCheck, err = CpfCheckUseCaseMock.UsecaseCpf.Check(*&newFakeCpf)
	assert.Nil(t, err)
	assert.NotEqual(t, cpfCheck, true)
}

func generateTrueCpf() []int {
	firstNine := generateFirstNineNumbers()
	cpf := append(firstNine, generateCpfFirstVerifierDigit(firstNine))
	cpf = append(cpf, generateCpfSecondVerifierDigit(cpf))
	fmt.Println("Cpf True: ", cpf)
	return cpf
}
func generateFakeCpf() []int {
	firstNine := generateFirstNineNumbers()
	cpf := append(firstNine, generateCpfFakeFirstVerifierDigit(firstNine))
	cpf = append(cpf, generateCpfFakeSecondVerifierDigit(cpf))
	fmt.Println("Cpf Fake: ", cpf)
	return cpf
}

func generateFirstNineNumbers() []int {
	var cpfFirstNine []int
	for len(cpfFirstNine) < 9 {
		newNumber := rand.Intn(9)
		cpfFirstNine = append(cpfFirstNine, newNumber)
		fmt.Println("Novo Array: ", cpfFirstNine)
	}
	if allElementsEqual(cpfFirstNine) {
		cpfFirstNine = generateFirstNineNumbers()
	}
	return cpfFirstNine
}

func generateCpfFirstVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (10 - i))
	}
	verifyDigit = (verifyDigit * 10) % 11
	return verifyDigit
}

func generateCpfSecondVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (11 - i))
	}
	verifyDigit = (verifyDigit * 10) % 11
	return verifyDigit
}

func generateCpfFakeFirstVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (10 - i))
	}
	verifyDigit = (verifyDigit * 10) % 11
	return verifyDigit + 1
}

func generateCpfFakeSecondVerifierDigit(fd []int) int {
	var verifyDigit int
	for i, num := range fd {
		verifyDigit = verifyDigit + (num * (11 - i))
	}
	verifyDigit = (verifyDigit * 10) % 11
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
